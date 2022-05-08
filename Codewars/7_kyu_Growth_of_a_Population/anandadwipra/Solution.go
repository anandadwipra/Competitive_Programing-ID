//  SPDX-License-Identifier: GPL-2.0-only
//
//  Copyright (C) 2021  anandadwipra <anandabiru04@gmail.com> 

package kata
func NbYear(p0 int, percent float64, aug int, p int) int {
  var i int
  for i= 0; p > p0 ;i++{
    p0 = p0 + int(float64(p0)*(percent/100.0)) +aug
  } 
  return i
}
