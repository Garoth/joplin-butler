# Joplin Butler: CLI Joplin Automation Tools

WARNING: Alpha version of the codebase

This provides a convenient wrapper around the Joplin WebClipper/Data Service:

    https://joplinapp.org/api/references/rest_api/

Motivation: the Joplin terminal app requires full sync, so it's not very
convenient for quick Joplin scripting. Joplin has a more convenient API
that it provides via a local JSON REST server, but using it via CURL
is kind of a hassle. This library provides user-friendly CLI scripting
tools for Joplin Desktop

## Example

Here we read the first two notes in the note list, and download detailed
versions of the notes that show metadata and content

    > joplin-butler notes | head -n2 | cut -d" " -f1 | \
        while read noteID; do  \
            joplin-butler note $noteID; |
        done

    ID: 1b16f5bfd4854f1f98345ef3d5dcc8c3
    Title: Some Best Insults
    ParentID: 269afae798c445fda34a2817cb0a96ea
    CreatedTime: 1307227257000
    UpdatedTime: 1307227414000
    Source: evernote
    ---
    - "Some cause happiness wherever they go; others, whenever they go." - Oscar Wilde
    - I'd agree with you but then we would both be wrong
    - If you were any less intelligent you would have to be watered twice a day.
    - Your IQ does not even make a respectable earthquake.
    - Somewhere a tree is making oxygen for you. You should apologize to it.

    ID: 6b2795f324154e33b3d2088625f44c66
    Title: Localization and Plurals - MDN
    ParentID: de70a04660784689a2b6cf32316e3632
    CreatedTime: 1318376811000
    UpdatedTime: 1318376813000
    Source: evernote.web.clip
    ---
    Introduced in Gecko 1.9
    (Firefox 3)

    You're likely here because you're localizing a .properties file and it had a link to this page. This page is to
     help explain how to localize these strings so that the correct plural form is shown to the user. E.g., "1 page
    " vs "2 pages".

    If you're here to make your code (e.g., extensions) localizable for plural forms, you can jump straight to [Dev
    eloping with PluralForm](https://developer.mozilla.org/en/#Developing_with_PluralForm), but you'll likely need 
    to localize the initial strings for your code, so it would be good to read through at least the Usage section a
    s well.

## Current Status

 * Automated authentication flow that pops up a dialog in Joplin
 * Configuration is saved to `~/.joplin-butler/config.json`
 * Can fetch the all-notes list and inspect individual notes

## TODO

 * First, implementing low-level raw access APIs for REST
 * Second, implementing higher level abstractions and conveniences
 * Third, providing useful high level scripts
