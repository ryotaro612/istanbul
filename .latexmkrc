#!/usr/bin/env perl
$lualatex = 'lualatex %O -synctex=1 -halt-on-error -shell-escape -interaction=nonstopmode %S';
$biber = 'biber %O --bblencoding=utf8 -u -U --output_safechars %B';
$bibtex = 'biblatex %O %B';
