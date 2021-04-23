# contents/pubcode_E.db/PublicationViewItem

|cid|name|type|notnull|dflt_value|pk|
| - | -- | -- | ----- | -------- | - |
|0|PublicationViewItemId|INTEGER|0||1|
|1|PublicationViewId|INTEGER|0||0|
|2|ParentPublicationViewItemId|INTEGER|0||0|
|3|Title|TEXT|0||0|
|4|SchemaType|INTEGER|0||0|
|5|ChildTemplateSchemaType|INTEGER|0||0|
|6|DefaultDocumentId|INTEGER|0||0|

| PublicationViewItemId | PublicationViewId | ParentPublicationViewItemId | Title | SchemaType | ChildTemplateSchemaType | DefaultDocumentId |
| - | - | - | - | - | - | - |
|1|1|-1|Good News|0|0|-1|
|2|1|1|Good News From God!|0||0|
|3|1|1|Title Page/Publishers’ Page|0||1|
|4|1|1|Table of Contents|0||2|
|5|1|1|How to Benefit From This Brochure|0||3|
