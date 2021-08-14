package libjw

// Most of this file is inspired by
// https://github.com/Miaosi001/JW-Library-macOS/blob/main/JWLibrary/Utility/JWPubExtractor.swift

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite driver

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
)

// THIS DOESN'T WORK
// DO NOT USE

type JWPUBWordMap struct {
	WordID                int
	Word                  string
	SearchIndexDocumentID int
	TextUnitCount         int
	WordOccurrenceCount   int
	TextUnitIndices       []byte
	PositionalList        []byte
	PositionalListIndex   []byte
}

// NOTE: This function have a lot of hardcoded values
// It is *not* ready for production usage
func JWPUBtoMarkdown(jwpub string) {
	//var wadd = make(map[string]int)
	path := helpers.GetDataDir() + "/_tmp_jwpub"
	log.Println(jwpub)
	err := helpers.Unzip(jwpub, path)
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.Unzip(path+"/contents", path+"/c")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("sqlite3", path+"/c/w_E_202110.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row, err := db.Query("SELECT WordId, Word FROM Word WHERE 1")
	if err != nil {
		log.Fatal(err)
	}
	var wordsmap []JWPUBWordMap
	for row.Next() {
		var wid int
		var w string
		err = row.Scan(&wid, &w)
		if err != nil {
			log.Fatal(err)
		}
		r := db.QueryRow("SELECT TextUnitIndices, PositionalList, PositionalListIndex FROM SearchIndexDocument WHERE WordId=?", wid)
		var tui []byte
		var pl []byte
		var pli []byte
		err = r.Scan(&tui, &pl, &pli)
		if err != nil {
			log.Fatal(err)
		}
		wordsmap = append(wordsmap, JWPUBWordMap{
			WordID:              wid,
			Word:                w,
			TextUnitIndices:     tui,
			PositionalList:      pl,
			PositionalListIndex: pli,
		})
	}

	sIndexes := wordsmap
	var loop = true
	var docID = 0
	var curDocIndex = []byte{128}
	var fullText = make(map[int]string)

	for loop {
		var finded = false
		for i := range sIndexes {
			if byteStartsWith(sIndexes[i].TextUnitIndices, []byte{128}) {
				if byteStartsWith(sIndexes[i].PositionalList, curDocIndex) {
					var rem = sIndexes[i].PositionalListIndex[0]
					if rem > 128 {
						finded = true
						var wd = sIndexes[i].Word
						//if wd != String(fullText[docID]?.split(separator: " ").last ?? "").unaccent() {
						//	print(curDocIndex, wd)
						//	fullText[docID]!.append(wd + " ")
						//}
						fullText[docID] += " " + wd
						sIndexes[i].PositionalList = sIndexes[i].PositionalList[len(curDocIndex):]
						//sIndexes[i].PositionalList = sIndexes[i].PositionalList.trimmingCharacters(in: .whitespacesAndNewlines)
						rem = rem - 1
						sIndexes[i].PositionalListIndex = sIndexes[i].PositionalListIndex[1:]
						sIndexes[i].PositionalListIndex = insertbyte(sIndexes[i].PositionalListIndex, rem, 0)
						//sIndexes[i].PositionalListIndex = rem + sIndexes[i].PositionalListIndex
						var curDocIndexArray = curDocIndex
						var repo = false
						for j := range curDocIndexArray {
							if j == 0 {
								if (curDocIndexArray[j] == 255 && len(curDocIndexArray) == 1) || (curDocIndexArray[j] == 127 && len(curDocIndexArray) > 1) {
									repo = true
									curDocIndex = []byte{0}
									if repo && j == len(curDocIndexArray)-1 {
										curDocIndex = append(curDocIndex, 129)
										repo = false
									}
								} else {
									curDocIndex = []byte{curDocIndexArray[j] + 1}
									repo = false
								}
							} else {
								if repo {
									if curDocIndexArray[j] == 255 {
										repo = true
										curDocIndex = append(curDocIndex, 129)
										if repo && j == len(curDocIndexArray)-1 {
											curDocIndex = append(curDocIndex, 129)
											repo = false
										}
									} else {
										curDocIndex = append(curDocIndex, curDocIndexArray[j]+1)
										repo = false
									}
								} else {
									curDocIndex = append(curDocIndex, curDocIndexArray[j])
								}
							}
						}
						break
					}
				}
			}
		}
		if !finded {
			var toRem []int
			for i := range sIndexes {
				var docI = sIndexes[i].TextUnitIndices[0]
				sIndexes[i].TextUnitIndices = sIndexes[i].TextUnitIndices[1:]
				if docI == 128 {
					//sIndexes[i].TextUnitIndices = sIndexes[i].TextUnitIndices.trimmingCharacters(in: .whitespacesAndNewlines)
					if len(sIndexes[i].TextUnitIndices) != 0 {
						docI = sIndexes[i].TextUnitIndices[0]
						sIndexes[i].TextUnitIndices = sIndexes[i].TextUnitIndices[1:]
						docI = docI - 1
						sIndexes[i].TextUnitIndices = insertbyte(sIndexes[i].TextUnitIndices, docI, 0)
					}
				} else {
					docI = docI - 1
					sIndexes[i].TextUnitIndices = insertbyte(sIndexes[i].TextUnitIndices, docI, 0)
				}
				if len(sIndexes[i].TextUnitIndices) == 0 {
					toRem = append(toRem, i)
				}
				var rem = sIndexes[i].PositionalListIndex[0]
				if rem == 128 {
					sIndexes[i].PositionalListIndex = sIndexes[i].PositionalListIndex[1:]
					//sIndexes[i].PositionalListIndex = sIndexes[i].PositionalListIndex.trimmingCharacters(in: .whitespacesAndNewlines)
				}
			}
			for i := len(toRem) - 1; i >= 0; i-- {
				sIndexes = append(sIndexes[:toRem[i]], sIndexes[toRem[i]+1:]...)
			}
			fmt.Println(fullText[docID])
			docID += 1
			curDocIndex = []byte{128}
		}
		if len(sIndexes) == 0 {
			loop = false
		}
	}
	//print(fullText)
	//for (id, text) in fullText where text != "" {
	//	let dir = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0].appendingPathComponent("w_I_202110/contents/\(id).txt")
	//	do {
	//		print(dir)
	//		try text.write(to: dir, atomically: true, encoding: String.Encoding.utf8)
	//	} catch {
	//		print("Error")
	//	}
	//}

}

func insertbyte(a []byte, c byte, i int) []byte {
	return append(a[:i], append([]byte{c}, a[i:]...)...)
}

func byteStartsWith(bs []byte, with []byte) bool {
	if len(bs) < len(with) {
		return false
	}
	for i := range with {
		if bs[i] != with[i] {
			return false
		}
	}
	return true
}
