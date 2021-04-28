# Converting meps_language_id to langcode

So in ~/LibJWgo/catalog/Publication.json we have ALL the publications.

Let's take this as an example:

```json
{
    "publication_root_key_id": 428,
    "meps_language_id": 198,
    "publication_type_id": 13,
    "issue_tag_number": 19991222,
    "title": "Przebudźcie się! — 1999",
    "issue_title": "Przebudźcie się!, 22 grudnia 1999",
    "short_title": "Przebudźcie się! — 1999",
    "cover_title": "",
    "undated_title": "Przebudźcie się!",
    "undated_reference_title": "Przebudźcie się!",
    "year": 1999,
    "symbol": "g99",
    "key_symbol": "g",
    "reserved": 0,
    "id": 340695
}
```

It is in Polish (lang code P, meps 198)

On the other hand we have English (lang code E, meps 0)

We also have this language endpoint, but it doesn't contain meps code.. so worthless for us.

I've mad some progress in `utils/getmepslangs`, for not it is usable, cd into that directory and run:

```bash
$ rm -rf /tmp/LibJWgo # If it's not your first try
$ go run main.go 2>../../libjw/mepsmap.go.tmp && mv ../../libjw/mepsmap.go.tmp ../../libjw/mepsmap.go
```

to update the map. After that you will notice some lines that do not fit there, simply delete them, then go lint and you are good to go