# JWPUB to Epub

Step-by-step guide to converting JWPUB, to EPUB.


# Blog - how things are going

Let's start with something simple, _[Good News From God!](https://www.jw.org/download/?output=html&pub=fg&fileformat=PDF%2CEPUB%2CJWPUB%2CRTF%2CTXT%2CBRL%2CBES%2CDAISY&alllangs=0&langwritten=E&txtCMSLang=P&isBible=0)_ looks nice.

It's 8mb in size, and it's a `zip`.

After unzipping we get:

```plain
fg_E.jwpub/
├── contents
└── manifest.json
```

 - [contents is a zip archive](#content)
 - [manifest.json is a formated JSON document](#manifest_json)

# /manifest.json

note: things after `//` got added

```json
{
  "name": "fg_E.jwpub", // This is file name codename_mepslangcode.format
  "hash": "3624ce30399a31e795bb02132081883adb3580aa59432990f57aa4ca47132a88", // sha256sum fg_E.jwpub/contents
  "timestamp": "2019-02-05T20:49:53Z",
  "version": 1,
  "expandedSize": 8990059,
  "contentFormat": "z-a",
  "htmlValidated": false,
  "mepsPlatformVersion": 2.1,
  "mepsBuildNumber": 8427,
  "publication": {
    "fileName": "fg_E.db",
    "type": 1,
    "title": "Good News From God!",
    "shortTitle": "Good News",
    "displayTitle": "Good News  (fg)",
    "referenceTitle": "Good News",
    "undatedReferenceTitle": "Good News",
    "symbol": "fg",
    "uniqueEnglishSymbol": "fg",
    "uniqueSymbol": "fg",
    "undatedSymbol": "fg",
    "englishSymbol": "fg",
    "language": 0,
    "hash": "b2014fb5b5008ee4e7802f76c9ed218208604cbd",
    "timestamp": "2019-02-05T20:49:53Z",
    "minPlatformVersion": 1,
    "schemaVersion": 8,
    "year": 2012,
    "issueId": 0,
    "issueNumber": 0,
    "variation": "",
    "publicationType": "Brochure",
    "rootSymbol": "fg",
    "rootYear": 2012,
    "rootLanguage": 0,
    "images": [
      {
        "signature": "2ed101473a1ec33cf45dbd6335e7cff7c5d700c0:600:771",
        "fileName": "1102012181_E_cvr.jpg",
        "type": "c",
        "attribute": "r",
        "width": 600,
        "height": 771
      },
      {
        "signature": "604fc195ebc7acea48155e1e5862b81e8862902b:120:120",
        "fileName": "1102012181_univ_sqr-120x120.jpg",
        "type": "t",
        "attribute": "r",
        "width": 120,
        "height": 120
      },
      {
        "signature": "604fc195ebc7acea48155e1e5862b81e8862902b:270:270",
        "fileName": "1102012181_univ_sqr-270x270.jpg",
        "type": "t",
        "attribute": "r",
        "width": 270,
        "height": 270
      },
      {
        "signature": "604fc195ebc7acea48155e1e5862b81e8862902b:600:600",
        "fileName": "1102012181_univ_sqr-600x600.jpg",
        "type": "t",
        "attribute": "r",
        "width": 600,
        "height": 600
      },
      {
        "signature": "a401a6f6630483bc33846bd46d320fbfa5215ba9:240:120",
        "fileName": "1102012181_univ_lsr-240x120.jpg",
        "type": "lsr",
        "attribute": "r",
        "width": 240,
        "height": 120
      },
      {
        "signature": "a401a6f6630483bc33846bd46d320fbfa5215ba9:560:280",
        "fileName": "1102012181_univ_lsr-560x280.jpg",
        "type": "lsr",
        "attribute": "r",
        "width": 560,
        "height": 280
      },
      {
        "signature": "a401a6f6630483bc33846bd46d320fbfa5215ba9:1200:600",
        "fileName": "1102012181_univ_lsr-1200x600.jpg",
        "type": "lsr",
        "attribute": "r",
        "width": 1200,
        "height": 600
      }
    ],
    "categories": [
      "brch"
    ],
    "attributes": [],
    "issueAttributes": [],
    "issueProperties": {
      "title": "",
      "undatedTitle": "",
      "coverTitle": "",
      "symbol": "",
      "undatedSymbol": ""
    }
  }
}
```

# /content

After unzipping it we got:

```plain
fg_E.jwpub/
├── contents
│   ├── 1102012181_E_cvr.jpg
│   ├── 1102012181_univ_lsr-1200x600.jpg
│   ├── 1102012181_univ_lsr-240x120.jpg
│   ├── 1102012181_univ_lsr-560x280.jpg
│   ├── 1102012181_univ_sqr-120x120.jpg
│   ├── 1102012181_univ_sqr-270x270.jpg
│   ├── 1102012181_univ_sqr-600x600.jpg
│   ├── 1102012181_univ_sqr.jpg
│   ├── 1102012182_E_cnt_2.jpg
│   ├── 1102012182_univ_cnt_1.jpg
│   ├── 1102012182_univ_sqr.jpg
│   ├── 1102012183_univ_cnt_1.jpg
│   ├── 1102012183_univ_sqr.jpg
│   ├── 1102012184_univ_cnt_1.jpg
│   ├── 1102012184_univ_cnt_2.jpg
│   ├── 1102012184_univ_cnt_3.jpg
│   ├── 1102012184_univ_sqr.jpg
│   ├── 1102012185_univ_cnt_1.jpg
│   ├── 1102012185_univ_cnt_2.jpg
│   ├── 1102012185_univ_sqr.jpg
│   ├── 1102012186_univ_cnt_1.jpg
│   ├── 1102012186_univ_cnt_2.jpg
│   ├── 1102012186_univ_sqr.jpg
│   ├── 1102012187_univ_cnt_1.jpg
│   ├── 1102012187_univ_cnt_2.jpg
│   ├── 1102012187_univ_cnt_3.jpg
│   ├── 1102012187_univ_sqr.jpg
│   ├── 1102012188_univ_cnt_1.jpg
│   ├── 1102012188_univ_cnt_2.jpg
│   ├── 1102012188_univ_cnt_3.jpg
│   ├── 1102012188_univ_sqr.jpg
│   ├── 1102012189_univ_cnt_1.jpg
│   ├── 1102012189_univ_cnt_2.jpg
│   ├── 1102012189_univ_cnt_3.jpg
│   ├── 1102012189_univ_sqr.jpg
│   ├── 1102012190_univ_cnt_1.jpg
│   ├── 1102012190_univ_cnt_2.jpg
│   ├── 1102012190_univ_sqr.jpg
│   ├── 1102012191_univ_cnt_1.jpg
│   ├── 1102012191_univ_cnt_2.jpg
│   ├── 1102012191_univ_cnt_3.jpg
│   ├── 1102012191_univ_cnt_4.jpg
│   ├── 1102012191_univ_sqr.jpg
│   ├── 1102012192_univ_cnt_1.jpg
│   ├── 1102012192_univ_cnt_2.jpg
│   ├── 1102012192_univ_cnt_3.jpg
│   ├── 1102012192_univ_sqr.jpg
│   ├── 1102012193_univ_cnt_1.jpg
│   ├── 1102012193_univ_cnt_2.jpg
│   ├── 1102012193_univ_cnt_3.jpg
│   ├── 1102012193_univ_cnt_4.jpg
│   ├── 1102012193_univ_sqr.jpg
│   ├── 1102012194_univ_cnt_1.jpg
│   ├── 1102012194_univ_cnt_2.jpg
│   ├── 1102012194_univ_sqr.jpg
│   ├── 1102012195_univ_cnt_1.jpg
│   ├── 1102012195_univ_cnt_2.jpg
│   ├── 1102012195_univ_sqr.jpg
│   ├── 1102012196_univ_cnt_1.jpg
│   ├── 1102012196_univ_cnt_2.jpg
│   ├── 1102012196_univ_cnt_3.jpg
│   ├── 1102012196_univ_sqr.jpg
│   ├── 1102012198_univ_cnt_1.jpg
│   ├── 1102012198_univ_cnt_2.jpg
│   ├── 1102012198_univ_sqr.jpg
│   ├── 502014331_univ_lsr.jpg
│   ├── 502015752_univ_wsr.jpg
│   ├── 502016852_univ_wsr.jpg
│   ├── 502016853_univ_wsr.jpg
│   ├── 502017850_univ_wsr.jpg
│   ├── 502017855_univ_wsr.jpg
│   ├── 502017858_univ_wsr.jpg
│   ├── 502018850_univ_wsr.jpg
│   ├── 502018853_univ_wsr.jpg
│   ├── 502018856_univ_wsr.jpg
│   └── fg_E.db
├── contents.orig
└── manifest.json

1 directory, 78 files
```

So a couple of jpegs, and a fg_E.db. I think that we all know how images work, so let's head over to [fg_E.db](#contents_fg_E_db)

# /contents/fg_E.db

(hidden info in source (a failed attempt))
<!-- 
This file is a sqlite3 database, here is a little overview of what is inside... It's a lot of stuff.
I've used this simple script to prepare the files in pubcode_db/ directory

```bash
for table in (sqlite3 fg_E.jwpub/contents/fg_E.db .tables | xargs | tr " " "\n")
    echo "# contents/pubcode_E.db/$table" > ../docs/jwpub/pubcode_db/"$table".md
    echo "" >> ../docs/jwpub/pubcode_db/"$table".md
    echo "|cid|name|type|notnull|dflt_value|pk|" >> ../docs/jwpub/pubcode_db/"$table".md
    echo "| - | -- | -- | ----- | -------- | - |" >> ../docs/jwpub/pubcode_db/"$table".md
    for i in (sqlite3 fg_E.jwpub/contents/fg_E.db "PRAGMA table_info($table)")
        echo "|$i|" >> ../docs/jwpub/pubcode_db/"$table".md
    end
    echo "" >> ../docs/jwpub/pubcode_db/"$table".md
    echo -n -e '|' >> ../docs/jwpub/pubcode_db/"$table".md 
    for i in (sqlite3 fg_E.jwpub/contents/fg_E.db "PRAGMA table_info($table)" | tr '|' ' ' | awk  '{print $2}')
        echo -n -e " $i |" >> ../docs/jwpub/pubcode_db/"$table".md
    end
    echo "" >> ../docs/jwpub/pubcode_db/"$table".md
    echo -n -e '|' >> ../docs/jwpub/pubcode_db/"$table".md
    for i in (sqlite3 fg_E.jwpub/contents/fg_E.db "PRAGMA table_info($table)" | tr '|' ' ' | awk  '{print $2}')
        echo -n -e " - |" >> ../docs/jwpub/pubcode_db/"$table".md
    end
    echo "" >> ../docs/jwpub/pubcode_db/"$table".md
    for j in (sqlite3 fg_E.jwpub/contents/fg_E.db "PRAGMA table_info($table)" | tr '|' ' ' | awk  '{print $2}')
        set i (sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT $j FROM $table LIMIT 1;")
        echo -n -e "|$i|" >> ../docs/jwpub/pubcode_db/"$table".md
    end
    echo "" >> ../docs/jwpub/pubcode_db/"$table".md
end
```
-->

