//  SPDX-License-Identifier: GPL-2.0-only
//
//  Copyright (C) 2021  anandadwipra <anandabiru04@gmail.com> 


package kata
import "strings"
func NoSpace(word string) string {
  return strings.ReplaceAll(word, " ", "")
}
