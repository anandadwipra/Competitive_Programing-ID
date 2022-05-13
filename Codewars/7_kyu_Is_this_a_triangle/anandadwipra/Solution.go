//  SPDX-License-Identifier: GPL-2.0-only
//
//  Copyright (C) 2021  anandadwipra <anandabiru04@gmail.com> 

package kata
func IsTriangle(a, b, c int) bool {
  if a+b>c && a+c>b && b+c>a{
  return true
  }else{
    return false
  }
}
