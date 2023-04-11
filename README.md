# Joplin Butler: CLI Joplin Automation Tools

<img width="400" src="https://github.com/Garoth/joplin-butler/blob/master/docs/img/note-butler-ai?raw=true" />
(Yes it's Midjourney art of a note taking butler :p)

WARNING: Beta version of the codebase

This provides a convenient wrapper around the Joplin WebClipper/Data Service:
https://joplinapp.org/api/references/rest_api/

Motivation: the Joplin terminal app requires full sync, so it's not very
convenient for quick Joplin scripting. Joplin has a more convenient API
that it provides via a local JSON REST server, but using it via CURL
is kind of a hassle. This library provides user-friendly CLI scripting
tools for Joplin Desktop

## Example

Create a note with title, resource, notebook (parent folder), and markdown body:

    > joplin-butler create resource ~/Downloads/Big\ airship/NE/still/still\ only0003.png

    ID: b8a8f3b5249943ca95aa0571bd905766
    Title: still only0003.png
    Filename: 
    CreatedTime: 1681174988187
    UpdatedTime: 1681174988187
    UserCreatedTime: 1681174988187
    UserUpdatedTime: 1681174988187
    Size: 1068527
    IsShared: 0

    > joplin-butler get folders

    269afae798c445fda34a2817cb0a96ea AE Misc
    4216484946714902922c15b075d39e28 Frequent
    dc80329d7b694117850e21556d59ddc7 Huge Lung
    029466b3d52f4cf2bf3aaab11ade8e40 Ideas
    a95cab228a9d4f118d4ad2c0a1b1adc5 Travel

    > joplin-butler create note 'My CLI Joplin Dream' \
        -body 'Begins today :D\n\n ![airship](:/b8a8f3b5249943ca95aa0571bd905766)' \
        -parent '269afae798c445fda34a2817cb0a96ea'

    ID: acd2cc52a6444e42a163fcefe94464d7
    Title: My CLI Joplin Dream
    ParentID: 1d4c4899bc7d4df4ad45cd575ac30043
    CreatedTime: 1681175174105
    UpdatedTime: 1681175174105
    Source: joplin-desktop
    ---
    Begins today :D

     ![airship](:/b8a8f3b5249943ca95aa0571bd905766)

The above code created the following note in Joplin Desktop instantly:

<img width="400" src="https://github.com/Garoth/joplin-butler/blob/master/docs/img/demo-note-1.png?raw=true" />

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

    ID: 6b2795f324154e33b3d2088625f44c66
    Title: Localization and Plurals - MDN
    ParentID: de70a04660784689a2b6cf32316e3632
    CreatedTime: 1318376811000
    UpdatedTime: 1318376813000
    Source: evernote.web.clip
    ---
    Introduced in Gecko 1.9
    (Firefox 3)

    You're likely here because you're localizing a .properties file and it had
    a link to this page. This page is to help explain how to localize these
    strings so that the correct plural form is shown to the user. E.g., "1 page"
    vs "2 pages" ...


Delete a note:

    > joplin-butler delete note/50fe5cf5dc6f4070a11c869cd4ec0138

    Successfully deleted notes/50fe5cf5dc6f4070a11c869cd4ec0138

## Development

The following commands put Go into the correct modules mode:

    export GO111MODULE=on
    export GOPATH=

Then `go run main.go` should work

## Current Status

 * Automated authentication flow that pops up a dialog in Joplin
 * Configuration is saved to `~/.joplin-butler/config.json`
 * Can fetch the all-notes list and inspect individual notes
 * Can create items like notes, folders, tags
 * Can delete items like notes, folders, tags
 * Can create resources / upload files by name

## TODO

 * Filters for which metadata to get about items
 * Attaching,Removing: attachments, tags, etc
 * Editing items, notes
