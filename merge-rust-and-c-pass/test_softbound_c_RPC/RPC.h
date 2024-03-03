/***************************************************************************
 **                                  _   _ ____  _
 **  Project                     ___| | | |  _ \| |
 **                             / __| | | | |_) | |
 **                            | (__| |_| |  _ <| |___
 **                             \___|\___/|_| \_\_____|
 **
 ** Copyright (C) 1998 - 2015, Daniel Stenberg, <daniel@haxx.se>, et al.
 **
 ** This software is licensed as described in the file COPYING, which
 ** you should have received as part of this distribution. The terms
 ** are also available at https://curl.haxx.se/docs/copyright.html.
 **
 ** You may opt to use, copy, modify, merge, publish, distribute and/or sell
 ** copies of the Software, and permit persons to whom the Software is
 ** furnished to do so, under the terms of the COPYING file.
 **
 ** This software is distributed on an "AS IS" basis, WITHOUT WARRANTY OF ANY
 ** KIND, either express or implied.
 **
 ****************************************************************************/ 
/* <DESC>
 ** simple HTTP POST using the easy interface
 ** </DESC>
 **/ 
#include <stdio.h>
#include <curl/curl.h>
#include <stdlib.h>
#include <string.h>

struct Output {
  char* buf;
};

static size_t get_output(void *buffer, size_t size, size_t nmemb, void *stream);
int call_another_func(char* func_name, char* input, char** output);

