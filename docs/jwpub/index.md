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

<!-- another fail
So.. 

`sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT Content FROM Extract WHERE ExtractId=1" | hexdump`

returned this:

```plain
0000000 c4a9 11c5 fdc5 c257 ad6d 94bf ea4a 4c1a
0000010 dc77 aaf8 dae4 3f6c 31b5 e68f cbf5 71a6
0000020 2cc8 047c fe48 a28e f7ed 446a ca28 e8a6
0000030 e076 2cdc f5ef 7bbf 32e8 6c2b d115 af18
0000040 29d4 1cbc 5780 af4a 07d9 ccfa 6932 01f1
0000050 28e0 1d2f 501e 4f61 6437 d713 d5ca e7c4
0000060 b842 26a2 a5d4 52f4 f9a9 8e25 7da5 8439
0000070 1b93 d1f3 ac1d ae2a c84a 3955 8f1d 0dcd
0000080 ece6 bb31 6fd4 4b02 e97a 991f 80e3 390b
0000090 1811 34bf dd31 c837 460c 68ad 0761 f028
00000a0 6d91 3d86 d7a1 240f 3b3d 659d d823 2e33
00000b0 5b3e 2666 dc05 13ce b236 c114 d3a1 64df
00000c0 91d8 2343 d73f 6650 e0fd 16c1 0ebb 7f51
00000d0 bee2 5b0e 43ae 80fa afa2 9f0f 8f0d e81c
00000e0 27ce 3e7e a978 44f0 3eeb 057c ed58 7334
00000f0 1616 05ef 2982 5f95 14ff ff23 1c5b 613d
0000100 8eed 5344 acab d2f5 4ee5 1f2c c883 0ca0
0000110 fd6e ae6f c32b 94be 36e8 e4bf da37 196c
0000120 e4f1 cc18 188a cdfa fc7d 5014 df01 6a5f
0000130 9bbf 75ad 121f 440a fe7a 74ed 6c73 8ecd
0000140 c23e 1a46 266d 40c4 0d66 831c 96fa 367b
0000150 98ff 9c11 e767 c9d5 dd0f e459 7e3f c2c9
0000160 2ac2 89e8 e28d 38ba d85a 3915 0566 ce0f
0000170 c5bf 1a97 7eb4 fcb0 1c3b eb2d 39ea 125f
0000180 63c3 2e0e cea0 f10e 42d0 4edd be39 a08a
0000190 58dd 4bdb 60ba 59cb a77b 359a 996a 76cd
00001a0 c079 61a6 4646 7a3c 808b 0a9e          
00001ac
```

(without hexdump it's just a piece of data)

And in same row we have `Caption` column with this inside

```html
<span class="eloc">bh pp. 18-26</span> <span class="etitle">The Bible—A Book From God</span>
```

So we know that it is a quotation from this some other book. Let's skip this (for now).

So let's `sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT Content FROM Document WHERE DocumentId=0" | hexdump`

```plain
0000000 35e3 d828 e024 1f89 c420 b63a 7215 760f
0000010 8966 64e6 9a28 88b8 34fe d1d0 4888 582e
0000020 200b d525 cd5b b113 eac2 b58a 3954 bfbd
0000030 1a2e e189 7aba 85b6 e06d cbe5 afa5 48bd
0000040 0562 2ad9 9c29 a57a b54e 4109 95f7 552e
0000050 f6e7 ca80 7604 85df 6224 6af8 6b98 e314
0000060 8faf 4878 f479 f838 eb3f 9fd0 489f d9de
0000070 f1fc f71f bd39 d75b 9ba6 cec3 9cd6 72a4
0000080 430e 2598 2176 ab8f e5b6 130d 3856 000a
000008f
```

Huh It's quite short. `Title` is equal to `Good News From God!` Which ughh.. Seems like it's not the correct thing I'm looking for. Let's hop into `Good News From God!`, with DocumentId=4


```plain
0000000 56a4 cef1 f027 eb20 f8e0 0f9c 2cb3 976f
0000010 0994 f9e1 07ab cde9 1d6a 8c72 ada9 a459
0000020 6185 26fa ee9f 3b88 4815 f8a2 4aac 343f
0000030 4d83 69d0 b0a3 c102 4c0e fc96 4c69 40d6
0000040 400f df94 e5bf 529b f347 e9ec 731d e9df
0000050 3d91 44ef 9fd7 274c 75db 22bd f0dd f4e9
0000060 123a 3a1c 4d0a b258 9e16 956d ae0b 56c8
0000070 f1af 4b49 9279 725c 7d23 1e63 be9b 8332
0000080 d203 1f58 488a f8a1 96cb 2fe3 0daf d388
0000090 6d7c 1e86 49cc ef77 389a 546d 3653 8fee
00000a0 dd11 52ee 38fc 2f5f 7874 7ad8 d93f 3f32
00000b0 4fb2 cb31 2d61 ca69 9dd3 2f80 dbb8 9a72
00000c0 4dcd 5db5 3f09 b445 0121 4793 e71c 9d91
00000d0 c8cd f557 a9dd 25f3 0d20 5926 187e 426d
00000e0 d53a 5d23 798e c70f 622c b819 86fc 7215
00000f0 be0f d661 344a 0b7d a3f9 927e b684 2438
0000100 351a bda8 2b2f 55c2 247a e206 97b8 2f62
0000110 540f c5b0 1f02 02ac 3a58 8ead 6edb 5873
0000120 0bd8 e6cd 5af1 4bd8 8a1b de3b 47e8 6c37
0000130 a79a 4e07 4757 b361 1a16 e216 5ae8 b2ab
0000140 6d9c 1cc9 ed03 4894 80b2 6b21 67f4 0ccd
0000150 fa6a a984 7a50 8ac0 6409 c740 1de8 83ab
0000160 a433 678c 4809 fb07 120e b134 d366 f0bc
0000170 7939 c10f 2ae9 5633 2d91 0329 3d5d 207e
0000180 ebd7 7d27 08c1 5880 86af 32db a476 1c44
0000190 8514 9b2e 6ddc 1ad4 caca b016 2a7e f14b
00001a0 c643 3e09 10e4 cdb2 2817 9569 0947 10a9
00001b0 dcfc acff ccf6 7377 f774 7933 bd19 3687
00001c0 ec9b 27ef 158a d577 33b4 ad73 ee2c f59a
00001d0 86f5 56ee dd7c b783 58e1 ee10 a82c 575e
00001e0 924c 4970 04ab 1871 1391 b2ee 29d8 662b
00001f0 22aa acfe 2d3b f6ec 4f32 883c 28c9 20a0
0000200 380b 0bf9 823f 0d6e 522c 2c32 69af 6495
0000210 3265 4b97 df8f b923 1e98 2a96 0cc0 f018
0000220 a5af fb02 a6c6 b966 1548 57bc 931e 6e5c
0000230 86a5 3230 a0b2 a025 f90a aacf c7c7 206b
0000240 9773 88ab ac3f 9aa7 74c3 f7db 0573 9ec5
0000250 1e46 983d 12c4 9870 9cdc 4952 7e45 187d
0000260 5b67 96f6 cde7 450a 000a               
0000269
```

Ah this one is longer.

This is how it should look like:

1\. What is the news from God?
------------------------------

[![People enjoying life on earth](https://assetsnffrgf-a.akamaihd.net/assets/m/1102012183/univ/art/1102012183_univ_cnt_1_md.jpg)](https://assetsnffrgf-a.akamaihd.net/assets/m/1102012183/univ/art/1102012183_univ_cnt_1_xl.jpg)

God wants people to enjoy life on earth. He created the earth and everything on it because he loves mankind. Soon he will act to provide a better future for people in every land. He will relieve mankind of the causes of suffering.​—_Read_ [_Jeremiah 29:11_](/en/library/bible/study-bible/books/jeremiah/29/#v24029011)_._

No government has ever succeeded in eliminating violence, disease, or death. But there is good news. Shortly, God will replace all human governments with his own government. Its subjects will enjoy peace and good health.​—_Read_ [_Isaiah 25:8;_](/en/library/bible/study-bible/books/isaiah/25/#v23025008) [_33:24;_](/en/library/bible/study-bible/books/isaiah/33/#v23033024) [_Daniel 2:44_](/en/library/bible/study-bible/books/daniel/2/#v27002044)_._

2\. Why is the good news urgent?
--------------------------------

Suffering will end only when God clears the earth of bad people. ([Zephaniah 2:3](/en/library/bible/study-bible/books/zephaniah/2/#v36002003)) When will that happen? God’s Word foretold the conditions that now threaten mankind. Current events indicate that God’s time to act is close.​—_Read_ [_2 Timothy 3:1-5_](/en/library/bible/study-bible/books/2-timothy/3/#v55003001-v55003005)_._

3\. What should we do?
----------------------

We should learn about God from his Word, the Bible. It is like a letter to us from a loving father. It tells us how to enjoy a better way of life now and how to enjoy everlasting life on earth in the future. True, some may not like it that you are receiving [help to understand the Bible](/en/bible-teachings/questions/understanding-the-bible/). But the opportunity of a better future is too good to miss.​—_Read_ [_Proverbs 29:25;_](/en/library/bible/study-bible/books/proverbs/29/#v20029025) [_Revelation 14:6, 7_](/en/library/bible/study-bible/books/revelation/14/#v66014006-v66014007)_._


Yup, I'm not anywhere close on finding it.

What I'm sure is that there are encoded:

 - Some heading/subheadings/fonts etc...
 - Images (probably by ID)
 - Links to other publications
 - And words, from the `Word` table.

For example 

`What is the news from God?` translates to:

|   |   |
| - | - |
| Decimal | `1246 616 1131 758 474 499` |
| Hex | 4de 268 46b 2f6 1da 1f3 |

Note: All words are in lower case.

-->

-------------------------------------------

Document.MepsDocumentId is equal to https://www.jw.org/finder?docid=thispart&prefer=lang&wtlocale=E

-------------------------

`$ sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT Content FROM Document WHERE DocumentId=8" | hexdump`
0000000 2fda 9755 0aa7 555b 92d4 9bbb a9db f1de
0000010 9daa 35e6 34ea 5170 672a 595e 6444 bdff
0000020 0a22                                   
0000022

wc --bytes returned 32

we also know that ContenLength returned: `7457` <!-- sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT Content FROM Document WHERE DocumentId=8" -->
And we also know that ParagraphCount is equal to: `20`

We also know that: ` sqlite3 fg_E.jwpub/contents/fg_E.db "SELECT ParagraphCount FROM Document WHERE DocumentId<8"`

2 + 12 + 18 + 12 + 10 + 21 + 19 + 19 = 113
so this is the first paragraph.

So following paragraphs are being included:

| DocumentParagraphId | DocumentId | ParagraphIndex | ParagraphNumberLabel | BeginPosition | EndPosition |
| ------------------- | ---------- | -------------- | -------------------- | ------------- | ----------- |
|114|8|1||10|159|
|115|8|2||161|246|
|116|8|3||319|398|
|117|8|4|1|674|1561|
|118|8|5|2|1563|2067|
|119|8|6|2|2069|2203|
|120|8|7||2386|2472|
|121|8|8|3|2496|2939|
|122|8|9|4|3172|3658|
|123|8|10||3713|3795|
|124|8|11|5|3819|4509|
|125|8|12||4564|4640|
|126|8|13|6|4664|5239|
|127|8|14|7|5241|5928|
|128|8|15||5983|6072|
|129|8|16|8|6504|7277|
|130|8|17||7331|7526|
|131|8|18||||
|132|8|19||6329|6460|
|133|8|20||||

So what do we know from that? No idea.