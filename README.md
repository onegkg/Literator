# Literator
## It's called literator because the trans is implied 🏳️‍⚧️

Literator is a Hebrew -> English transliteration engine written in Go. It's simultaneously a semi-useful project for me and also an excuse for me to learn Go.

The goal is for this project to serve as the backend for an online transliterator, but we're just focusing on making it work at the moment.

This repo is just a snapshot of my current work on the project. At any given point, until otherwise stated, there is no guarantee that this project will work, or even compile.

### TODO:
 - Testing (def integration tests, maybe unit tests too?) 
 - Create punctuation list
 - Add doc comments
 - Refactor (get rid of types folder)
 - Implement Edgecases
 - Actually fix transliteration
  - Furtive Patach
  - No shva at end of word
  - Recognize when alephs, ayins, and yuds are silent
  - Shva Na vs. Shva Nach
  - Psik Hei and silent hei
  - Edge case for yerushalayim
 - Serve over http
