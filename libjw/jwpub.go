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

	var loop = true
	var docID = 0
	var curDocIndex = []byte{128}
	var fullText = make(map[int]string, 255)

	sIndexes := wordsmap
	for loop {

		var finded = false
		for i := range sIndexes {
			if sIndexes[i].WordID == 123 {
				log.Println(sIndexes[i].Word, sIndexes[i].TextUnitIndices, byteStartsWith(sIndexes[i].TextUnitIndices, []byte{128}), sIndexes[i].PositionalList, byteStartsWith(sIndexes[i].PositionalList, curDocIndex), curDocIndex)
			}
			if byteStartsWith(sIndexes[i].TextUnitIndices, []byte{128}) {
				if byteStartsWith(sIndexes[i].PositionalList, curDocIndex) {
					var rem = sIndexes[i].PositionalListIndex[0]
					if rem > 128 {
						finded = true
						wd := sIndexes[i].Word
						fullText[docID] += " " + wd
						sIndexes[i].PositionalList = sIndexes[i].PositionalList[len(curDocIndex):]
						rem = rem - 1
						sIndexes[i].PositionalListIndex[0] = rem
						curDocIndexArray := curDocIndex
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
					} else {
						sIndexes[i].PositionalListIndex = sIndexes[i].PositionalListIndex[1:]
					}
				}
			}
		}

		if fullText[docID] != "" {
			fmt.Println("fullText[docID:", docID, "]:", fullText[docID])
		}
		if !finded {
			var toRem []int = []int{}
			for i := range sIndexes {
				//var docI = sIndexes[i].TextUnitIndices.prefix(3)
				//sIndexes[i].TextUnitIndices.removeFirst(3)
				var docI byte = 0
				if len(sIndexes[i].TextUnitIndices) > 0 {
					docI = sIndexes[i].TextUnitIndices[0]
				}
				if len(sIndexes[i].TextUnitIndices) == 0 {
					toRem = append(toRem, i)
				} else {
					sIndexes[i].TextUnitIndices = sIndexes[i].TextUnitIndices[1:]
					if docI == 128 {
						if len(sIndexes[i].TextUnitIndices) != 0 {
							sIndexes[i].TextUnitIndices = insertbyte(sIndexes[i].TextUnitIndices, docI-1, 0)
						}
					} else {
						docI--
						sIndexes[i].TextUnitIndices = insertbyte(sIndexes[i].TextUnitIndices, docI, 0)
					}
					if len(sIndexes[i].PositionalListIndex) > 0 && sIndexes[i].PositionalListIndex[0] == 128 {
						sIndexes[i].PositionalListIndex = sIndexes[i].PositionalListIndex[1:]
					}
				}
			}
			for i := len(toRem) - 1; i >= 0; i-- {
				log.Println(i, docID, "toRem2", sIndexes[toRem[i]].Word)
				sIndexes = append(sIndexes[:toRem[i]], sIndexes[toRem[i]+1:]...)
			}
			docID++
			curDocIndex = []byte{128}
		}
		if len(sIndexes) == 0 {
			loop = false
		} else {
			log.Println("len(sIndexes):", len(sIndexes))
		}
		if docID > 10000 {
			log.Fatal("docID > 10000, this should not happen.")
		}
	}
	//	for (id, text) in fullText where text != "" {
	//		let dir = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0].appendingPathComponent("w_I_202110/contents/\(id).txt")
	//		do {
	//			print(dir)
	//			try text.write(to: dir, atomically: true, encoding: String.Encoding.utf8)
	//		} catch {
	//			print("Error")
	//		}
	//	}

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
