package libjw

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.mrcyjanek.net/mrcyjanek/jwapi/helpers"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/jwhttp"
	"git.mrcyjanek.net/mrcyjanek/jwapi/libjw/structs"
	_ "github.com/mattn/go-sqlite3" // sqlite driver
)

// Publications - Store information about all Publications
var Publications = make(map[string]*structs.Publication)

// GetCatalog downloads catalog.db
func GetCatalog(path string, onlyifrequired bool) {
	if onlyifrequired {
		if string(helpers.Get("catalog_downloaded")) == "yes" {
			go func() {
				fmt.Println("[libjw][GetCatalog] Fetching version... (1)")
				version := jwhttp.CatalogsPublicationsV4Manifest().Current
				if version == string(helpers.Get("catalog_version")) {
					fmt.Println("[libjw][GetCatalog] Catalog is up to date (2)")
					return
				}
				var callbacks []AlertCallback
				callbacks = append(callbacks, AlertCallback{
					Title:    "Cancel",
					Endpoint: "/api/alerts/cancel/" + strconv.Itoa(len(Alerts)),
				})
				callbacks = append(callbacks, AlertCallback{
					Title:    "Update",
					Endpoint: "/api/updateCatalog",
				})
				CreateAlert(Alert{
					Title:       "New version of catalog is available!",
					Description: "Current version of publications catalog is '" + string(helpers.Get("catalog_version")) + "' while on JW server there is a version '" + version + "'" + ".\nThere is no need to update, unless you are waiting for new publication. After clicking update, a 50mb file will be downloaded, consider updating on a unmetred network.",
					Color:       "info",
					Cause:       "[libjw][GetCatalog]",
					Callbacks:   callbacks,
				})
			}()
			return
		} else {
			log.Println("[libjw][GetCatalog] - forcing update of catalog because.")
		}
	}
	fmt.Println("[libjw][GetCatalog] Fetching version...")
	version := jwhttp.CatalogsPublicationsV4Manifest().Current
	if version == string(helpers.Get("catalog_version")) {
		fmt.Println("[libjw][GetCatalog] Catalog is up to date")
		return
	}
	fmt.Println("[libjw][GetCatalog] version: " + version)
	alert := CreateAlert(Alert{
		Title:       "Updating catalog...",
		Description: "Version: " + version,
	})
	fmt.Println("[libjw][GetCatalog] Downloading gzipped catalog")
	UpdateDescription(alert, "Updating catalog...", "Downloading gzipped catalog")
	catalogBytesGz := jwhttp.CatalogsPublicationsV4CatalogDbGz(version)
	UpdateDescription(alert, "Updating catalog...", "Extracting...")
	fmt.Println("[libjw][GetCatalog] Extracting...")
	catalogBytesReader := bytes.NewReader(catalogBytesGz)
	reader, err := gzip.NewReader(catalogBytesReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	catalog, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[libjw][GetCatalog] Writing to " + path)
	UpdateDescription(alert, "Updating catalog...", "Writing...")
	ioutil.WriteFile(path, catalog, 0770)
	helpers.Set("catalog_version", []byte(version))
	DeleteAlert(alert)
	fmt.Println("[libjw][GetCatalog] Done")
	helpers.Mkdir(helpers.GetDataDir() + "/catalog")
	ParseCatalog(helpers.GetDataDir()+"/raw/catalog.db", helpers.GetDataDir()+"/catalog")
	helpers.Set("catalog_downloaded", []byte("yes"))
}

// ParseCatalog - convert sqlite3 catalog.db into few json files
func ParseCatalog(catalog string, target string) error {
	fmt.Println("[libjw][ParseCatalog] start")
	db, err := sql.Open("sqlite3", catalog)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	fmt.Println("[libjw][ParseCatalog] Setting null to default values.")

	// Start dumping sql
	var wg sync.WaitGroup

	wg.Add(1)

	// AvailableBibleBook
	wg.Add(1)
	go func() {
		var element = "AvailableBibleBook"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBAvailableBibleBook{}
		for row.Next() {
			var Book int
			var PublicationID int
			err = row.Scan(&Book, &PublicationID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBAvailableBibleBook{
				Book:          Book,
				PublicationID: PublicationID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// CuratedAsset
	wg.Add(1)
	go func() {
		var element = "CuratedAsset"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBCuratedAsset{}
		for row.Next() {
			var ListType int
			var SortOrder int
			var PublicationAssetID int
			err = row.Scan(&ListType, &SortOrder, &PublicationAssetID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBCuratedAsset{
				ListType:           ListType,
				SortOrder:          SortOrder,
				PublicationAssetID: PublicationAssetID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// DatedText
	wg.Add(1)
	go func() {
		var element = "DatedText"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBDatedText{}
		for row.Next() {
			var Class int
			var Start string
			var End string
			var PublicationID int
			row.Scan(&Class, &Start, &End, &PublicationID)
			silce = append(silce, structs.DBDatedText{
				Class:         Class,
				Start:         Start,
				End:           End,
				PublicationID: PublicationID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// ImageAsset
	wg.Add(1)
	go func() {
		var element = "ImageAsset"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBImageAsset{}
		for row.Next() {
			var Width int
			var Height int
			var NameFragment string
			var Size int
			var MimeType string
			var ID int
			err = row.Scan(&Width, &Height, &NameFragment, &Size, &MimeType, &ID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBImageAsset{
				Width:        Width,
				Height:       Height,
				NameFragment: NameFragment,
				Size:         Size,
				MimeType:     MimeType,
				ID:           ID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// Publication
	wg.Add(1)
	go func() {
		var element = "Publication"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT PublicationRootKeyId, MepsLanguageId, PublicationTypeId, IssueTagNumber, Title, IFNULL(IssueTitle, ''), ShortTitle, IFNULL(CoverTitle, ''), IFNULL(UndatedTitle, ''), IFNULL(UndatedReferenceTitle, ''), Year, Symbol, IFNULL(KeySymbol, ''), Reserved, Id FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublication{}
		//s, _ := row.Columns()
		//log.Println("publication, col", s)
		//PublicationRootKeyId MepsLanguageId PublicationTypeId IssueTagNumber Title IssueTitle ShortTitle CoverTitle UndatedTitle UndatedReferenceTitle Year Symbol KeySymbol Reserved Id
		for row.Next() {
			var PublicationRootKeyID int
			var MepsLanguageID int
			var PublicationTypeID int
			var IssueTagNumber int
			var Title string
			var IssueTitle string
			var ShortTitle string
			var CoverTitle string
			var UndatedTitle string
			var UndatedReferenceTitle string
			var Year int
			var Symbol string
			var KeySymbol string
			var Reserved int
			var ID int
			err = row.Scan(&PublicationRootKeyID, &MepsLanguageID, &PublicationTypeID, &IssueTagNumber, &Title, &IssueTitle, &ShortTitle, &CoverTitle, &UndatedTitle, &UndatedReferenceTitle, &Year, &Symbol, &KeySymbol, &Reserved, &ID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublication{
				PublicationRootKeyID:  PublicationRootKeyID,
				MepsLanguageID:        MepsLanguageID,
				PublicationTypeID:     PublicationTypeID,
				IssueTagNumber:        IssueTagNumber,
				Title:                 Title,
				IssueTitle:            IssueTitle,
				ShortTitle:            ShortTitle,
				CoverTitle:            CoverTitle,
				UndatedTitle:          UndatedTitle,
				UndatedReferenceTitle: UndatedReferenceTitle,
				Year:                  Year,
				Symbol:                Symbol,
				KeySymbol:             KeySymbol,
				Reserved:              Reserved,
				ID:                    ID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationAsset
	wg.Add(1)
	go func() {
		var element = "PublicationAsset"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT PublicationId, MepsLanguageId, Signature, Size, ExpandedSize, MimeType, SchemaVersion, MinPlatformVersion, CatalogedOn, LastUpdated, IFNULL(GenerallyAvailableDate, ''), IFNULL(ConventionReleaseDayNumber, 0), Id FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationAsset{}
		for row.Next() {
			var PublicationID int
			var MepsLanguageID int
			var Signature string
			var Size int
			var ExpandedSize int
			var MimeType string
			var SchemaVersion int
			var MinPlatformVersion int
			var CatalogedOn string
			var LastUpdated string
			var GenerallyAvailableDate string
			var ConventionReleaseDayNumber int
			var ID int
			err = row.Scan(&PublicationID, &MepsLanguageID, &Signature, &Size, &ExpandedSize, &MimeType, &SchemaVersion, &MinPlatformVersion, &CatalogedOn, &LastUpdated, &GenerallyAvailableDate, &ConventionReleaseDayNumber, &ID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationAsset{
				PublicationID:              PublicationID,
				MepsLanguageID:             MepsLanguageID,
				Signature:                  Signature,
				Size:                       Size,
				ExpandedSize:               ExpandedSize,
				MimeType:                   MimeType,
				SchemaVersion:              SchemaVersion,
				MinPlatformVersion:         MinPlatformVersion,
				CatalogedOn:                CatalogedOn,
				LastUpdated:                LastUpdated,
				GenerallyAvailableDate:     GenerallyAvailableDate,
				ConventionReleaseDayNumber: ConventionReleaseDayNumber,
				ID:                         ID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationAssetImageMap
	wg.Add(1)
	go func() {
		var element = "PublicationAssetImageMap"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationAssetImageMap{}
		for row.Next() {
			var PublicationAssetID int
			var ImageAssetID int
			err = row.Scan(&PublicationAssetID, &ImageAssetID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationAssetImageMap{
				PublicationAssetID: PublicationAssetID,
				ImageAssetID:       ImageAssetID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationAttribute
	wg.Add(1)
	go func() {
		var element = "PublicationAttribute"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationAttribute{}
		for row.Next() {
			var Name string
			var ID int
			err = row.Scan(&Name, &ID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationAttribute{
				Name: Name,
				ID:   ID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationAttributeMap
	wg.Add(1)
	go func() {
		var element = "PublicationAttributeMap"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationAttributeMap{}
		for row.Next() {
			var PublicationID int
			var PublicationAttributeID int
			err = row.Scan(&PublicationID, &PublicationAttributeID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationAttributeMap{
				PublicationID:          PublicationID,
				PublicationAttributeID: PublicationAttributeID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationDocument
	wg.Add(1)
	go func() {
		var element = "PublicationDocument"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationDocument{}
		for row.Next() {
			var DocumentID int
			var PublicationID int
			err = row.Scan(&DocumentID, &PublicationID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationDocument{
				DocumentID:    DocumentID,
				PublicationID: PublicationID,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// PublicationRootKey
	wg.Add(1)
	go func() {
		var element = "PublicationRootKey"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT Symbol, IFNULL(Year, 0), Language, Id FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		var silce = []structs.DBPublicationRootKey{}
		for row.Next() {
			var Symbol string
			var Year int
			var Language int
			var ID int
			err = row.Scan(&Symbol, &Year, &Language, &ID)
			if err != nil {
				log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
			}
			silce = append(silce, structs.DBPublicationRootKey{
				Symbol:   Symbol,
				Year:     Year,
				Language: Language,
			})
		}
		res, err := json.Marshal(silce)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	// Revision
	wg.Add(1)
	go func() {
		var element = "Revision"
		fmt.Println("[libjw][ParseCatalog] " + element + ":start")
		row, err := db.Query("SELECT * FROM " + element + " WHERE 1")
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+" - err", err)
		}
		row.Next()
		var Level int
		var Created string
		err = row.Scan(&Level, &Created)
		if err != nil {
			log.Fatal("[libjw][ParseCatalog] "+element+":error", err)
		}
		revision := structs.DBRevision{
			Level:   Level,
			Created: Created,
		}
		res, err := json.Marshal(revision)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(target+"/"+element+".json", res, 0770)
		fmt.Println("[libjw][ParseCatalog] " + element + ":end")
		wg.Done()
	}()

	wg.Done()
	wg.Wait()
	fmt.Println("[libjw][ParseCatalog] end")
	return nil
}

// BuildPublications fill Publications map with all publications
func BuildPublications(catalog string) {
	fmt.Println("[libjw][BuildPublication] start")
	// Create variables
	AvailableBibleBook := []structs.DBAvailableBibleBook{}
	CuratedAsset := []structs.DBCuratedAsset{}
	DatedText := []structs.DBDatedText{}
	ImageAsset := []structs.DBImageAsset{}
	PublicationAssetImageMap := []structs.DBPublicationAssetImageMap{}
	PublicationAsset := []structs.DBPublicationAsset{}
	PublicationAttribute := []structs.DBPublicationAttribute{}
	PublicationAttributeMap := []structs.DBPublicationAttributeMap{}
	PublicationDocument := []structs.DBPublicationDocument{}
	Publication := []structs.DBPublication{}               // Used
	PublicationRootKey := []structs.DBPublicationRootKey{} // Used
	Revision := structs.DBRevision{}

	// Read JSON
	readJSONFile(catalog+"/AvailableBibleBook.json", &AvailableBibleBook)
	readJSONFile(catalog+"/CuratedAsset.json", &CuratedAsset)
	readJSONFile(catalog+"/DatedText.json", &DatedText)
	readJSONFile(catalog+"/ImageAsset.json", &ImageAsset)
	readJSONFile(catalog+"/PublicationAssetImageMap.json", &PublicationAssetImageMap)
	readJSONFile(catalog+"/PublicationAsset.json", &PublicationAsset)
	readJSONFile(catalog+"/PublicationAttribute.json", &PublicationAttribute)
	readJSONFile(catalog+"/PublicationAttributeMap.json", &PublicationAttributeMap)
	readJSONFile(catalog+"/PublicationDocument.json", &PublicationDocument)
	readJSONFile(catalog+"/Publication.json", &Publication)               // Used
	readJSONFile(catalog+"/PublicationRootKey.json", &PublicationRootKey) // Used
	readJSONFile(catalog+"/Revision.json", &Revision)

	for i := range PublicationRootKey {
		pubr := PublicationRootKey[i]
		ld := []structs.LangData{}
		struc := structs.Publication{
			Code:     pubr.Symbol,
			Year:     pubr.Year,
			LangData: &ld,
		}
		Publications[pubr.Symbol] = &struc
	}
	fmt.Println("[libjw][BuildPublication] end")
}

// GetLanguages - return struct of all languages
func GetLanguages() structs.Languages {
	filename := helpers.GetDataDir() + "/languages.json"
	info, err := os.Stat(filename)
	lastMod := int64(0)
	if err != nil {
		fmt.Println("[libjw][GetLanguages] WARN! languages.json not found")
	} else {
		lastMod = time.Now().Unix() - info.ModTime().Unix()
	}
	if lastMod > 7*24*60*60 || err != nil {
		fmt.Println("[libjw][GetLanguages] Updating languages.json...")
		langs := jwhttp.SiteLanguages("en")
		res, err := json.Marshal(langs)
		if err != nil {
			fmt.Println(err)
			return structs.Languages{
				Status: 500,
			}
		}
		ioutil.WriteFile(filename, res, 0770)
		return langs
	}
	var langs structs.Languages
	readJSONFile(filename, &langs)
	return langs
}

func readJSONFile(file string, s interface{}) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = json.Unmarshal(data, s)
	if err != nil {
		fmt.Println(err)
	}
}

// GetPublication - Fetch publication and extract it, so it can be later displayed in webui
func GetPublication(publication string, language string, format string, issue string) (pub structs.PublicationV2, e error) {
	exp := issue
	if exp == "" {
		exp = "root"
	}

	reg, err := regexp.Compile(`[^a-zA-Z0-9\-]+`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	publication = reg.ReplaceAllString(publication, "")
	issue = reg.ReplaceAllString(issue, "")

	partpath := "/data/publications/" + publication + "/" + exp + "/" + language + "/"
	pubURI := ""
	if issue == "" {
		pubURI = "api/publications/" + publication + "/"
	} else {
		pubURI = "api/publications/" + publication + "_" + issue + "/"
	}
	extractpath := helpers.GetDataDir() + partpath
	_, err = os.Stat(extractpath)
	struc := structs.PublicationV2{
		Title:  "<undefined>",
		Format: format,
		Path:   partpath,
		URI:    pubURI,
	}
	if err != nil {
		format = strings.ToUpper(format)
		if publication == "" {
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] publication invalid")
		}
		if language == "" {
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] language invalid")
		}
		if format == "" {
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] format invalid")
		}
		fmt.Println("[libjw][GetPublication] Downloading data...", publication, language, format)
		pubinfo, status := jwhttp.ApisPubMediaGETPUBMEDIALINKS(language, publication, format, issue)
		if len(status) != 0 {
			switch status[0].Status {
			case 400:
				return structs.PublicationV2{}, errors.New("[libjw][GetPublication] Failed to download publication! It never existed")
			case 404:
				return structs.PublicationV2{}, errors.New("[libjw][GetPublication] Failed to download publication! It existed but got deleted from JW servers or isn't available in this format/language")
			}
		}
		pub := pubinfo.Files[language][format][0]
		helpers.Mkdir(extractpath)
		fmt.Println("[libjw][GetPublication] Downloading...", pub.Title)
		resp, err := http.Get(pub.File.URL)
		if err != nil {
			fmt.Println(err)
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] " + err.Error() + " (is publication/language/format valid?)")
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] " + err.Error() + " (probably network error)")
		}
		fmt.Println("[libjw][GetPublication] Writing...", pub.Title)
		f, err := ioutil.TempFile("", "tmpdlfile_*")
		if err != nil {
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] " + err.Error() + " (failed to create TempFile)")
		}
		f.Write(body)
		f.Sync()
		fmt.Println("[libjw][GetPublication] Extracting...", pub.Title)
		err = helpers.Unzip(f.Name(), extractpath)
		if err != nil {
			return structs.PublicationV2{}, errors.New("[libjw][GetPublication] " + err.Error() + " (zipslip of something? Maybe corrupted download, failed to uzip)")
		}
		struc = structs.PublicationV2{
			Title:  pub.Title,
			Format: format,
			Path:   partpath,
			URI:    pubURI,
		}
		res, err := json.Marshal(struc)
		if err != nil {
			fmt.Println(err)
			return structs.PublicationV2{
				Title: "FAILED",
			}, errors.New("[libjw][GetPublication] " + err.Error() + " (failed to marshal json)")
		}
		ioutil.WriteFile(helpers.GetDataDir()+struc.Path+"config.json", res, 0770)
		return struc, nil
	}
	readJSONFile(helpers.GetDataDir()+struc.Path+"config.json", &struc)
	return struc, nil
}

// DecodeContentToc - decode XML content.opf from extraced EPUB
func DecodeContentToc(path string) structs.EPUBTocNcx {
	var toc structs.EPUBTocNcx
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	err = xml.NewDecoder(file).Decode(&toc)
	if err != nil {
		fmt.Println(err)
		fmt.Println("[libjw][DecodeContentToc] I shouldn't reach this point")
		return structs.EPUBTocNcx{}
	}
	return toc
}

// DecodeContentOpf - decode XML content.opf from extraced EPUB
func DecodeContentOpf(path string) structs.EPUBContentOpf {
	var contentOpf structs.EPUBContentOpf
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	err = xml.NewDecoder(file).Decode(&contentOpf)
	if err != nil {
		fmt.Println(err)
		fmt.Println("[libjw][DecodeContentOpf] I shouldn't reach this point")
		return structs.EPUBContentOpf{}
	}
	return contentOpf
}
