#!/usr/bin/env perl


$latex = 'lualatex %O -synctex=1 -interaction=nonstopmode %S';
$biber = 'biber %O --bblencoding=utf8 -u -U --output_safechars %B';
$bibtex = 'biblatex %O %B';