package main

import (
	"log"
	"net/http"
	  _ "embed"
)

//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T000000Z.png
var img1 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T010000Z.png
var img2 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T020000Z.png
var img3 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T030000Z.png
var img4 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T040000Z.png
var img5 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T050000Z.png
var img6 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T060000Z.png
var img7 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T070000Z.png
var img8 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T080000Z.png
var img9 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T090000Z.png
var img10 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T100000Z.png
var img11 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T110000Z.png
var img12 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T120000Z.png
var img13 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T130000Z.png
var img14 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T140000Z.png
var img15 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T150000Z.png
var img16 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T160000Z.png
var img17 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T170000Z.png
var img18 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T180000Z.png
var img19 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T190000Z.png
var img20 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T200000Z.png
var img21 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T210000Z.png
var img22 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T220000Z.png
var img23 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950101T230000Z.png
var img24 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T000000Z.png
var img25 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T010000Z.png
var img26 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T020000Z.png
var img27 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T030000Z.png
var img28 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T040000Z.png
var img29 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T050000Z.png
var img30 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T060000Z.png
var img31 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T070000Z.png
var img32 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T080000Z.png
var img33 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T090000Z.png
var img34 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T100000Z.png
var img35 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T110000Z.png
var img36 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T120000Z.png
var img37 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T130000Z.png
var img38 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T140000Z.png
var img39 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T150000Z.png
var img40 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T160000Z.png
var img41 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T170000Z.png
var img42 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T180000Z.png
var img43 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T190000Z.png
var img44 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T200000Z.png
var img45 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T210000Z.png
var img46 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T220000Z.png
var img47 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950102T230000Z.png
var img48 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T000000Z.png
var img49 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T010000Z.png
var img50 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T020000Z.png
var img51 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T030000Z.png
var img52 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T040000Z.png
var img53 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T050000Z.png
var img54 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T060000Z.png
var img55 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T070000Z.png
var img56 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T080000Z.png
var img57 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T090000Z.png
var img58 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T100000Z.png
var img59 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T110000Z.png
var img60 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T120000Z.png
var img61 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T130000Z.png
var img62 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T140000Z.png
var img63 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T150000Z.png
var img64 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T160000Z.png
var img65 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T170000Z.png
var img66 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T180000Z.png
var img67 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T190000Z.png
var img68 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T200000Z.png
var img69 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T210000Z.png
var img70 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T220000Z.png
var img71 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950103T230000Z.png
var img72 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T000000Z.png
var img73 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T010000Z.png
var img74 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T020000Z.png
var img75 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T030000Z.png
var img76 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T040000Z.png
var img77 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T050000Z.png
var img78 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T060000Z.png
var img79 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T070000Z.png
var img80 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T080000Z.png
var img81 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T090000Z.png
var img82 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T100000Z.png
var img83 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T110000Z.png
var img84 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T120000Z.png
var img85 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T130000Z.png
var img86 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T140000Z.png
var img87 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T150000Z.png
var img88 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T160000Z.png
var img89 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T170000Z.png
var img90 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T180000Z.png
var img91 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T190000Z.png
var img92 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T200000Z.png
var img93 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T210000Z.png
var img94 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T220000Z.png
var img95 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950104T230000Z.png
var img96 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T000000Z.png
var img97 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T010000Z.png
var img98 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T020000Z.png
var img99 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T030000Z.png
var img100 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T040000Z.png
var img101 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T050000Z.png
var img102 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T060000Z.png
var img103 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T070000Z.png
var img104 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T080000Z.png
var img105 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T090000Z.png
var img106 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T100000Z.png
var img107 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T110000Z.png
var img108 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T120000Z.png
var img109 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T130000Z.png
var img110 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T140000Z.png
var img111 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T150000Z.png
var img112 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T160000Z.png
var img113 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T170000Z.png
var img114 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T180000Z.png
var img115 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T190000Z.png
var img116 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T200000Z.png
var img117 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T210000Z.png
var img118 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T220000Z.png
var img119 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950105T230000Z.png
var img120 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T000000Z.png
var img121 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T010000Z.png
var img122 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T020000Z.png
var img123 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T030000Z.png
var img124 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T040000Z.png
var img125 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T050000Z.png
var img126 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T060000Z.png
var img127 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T070000Z.png
var img128 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T080000Z.png
var img129 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T090000Z.png
var img130 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T100000Z.png
var img131 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T110000Z.png
var img132 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T120000Z.png
var img133 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T130000Z.png
var img134 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T140000Z.png
var img135 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T150000Z.png
var img136 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T160000Z.png
var img137 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T170000Z.png
var img138 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T180000Z.png
var img139 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T190000Z.png
var img140 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T200000Z.png
var img141 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T210000Z.png
var img142 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T220000Z.png
var img143 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950106T230000Z.png
var img144 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T000000Z.png
var img145 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T010000Z.png
var img146 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T020000Z.png
var img147 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T030000Z.png
var img148 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T040000Z.png
var img149 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T050000Z.png
var img150 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T060000Z.png
var img151 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T070000Z.png
var img152 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T080000Z.png
var img153 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T090000Z.png
var img154 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T100000Z.png
var img155 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T110000Z.png
var img156 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T120000Z.png
var img157 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T130000Z.png
var img158 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T140000Z.png
var img159 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T150000Z.png
var img160 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T160000Z.png
var img161 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T170000Z.png
var img162 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T180000Z.png
var img163 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T190000Z.png
var img164 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T200000Z.png
var img165 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T210000Z.png
var img166 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T220000Z.png
var img167 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950107T230000Z.png
var img168 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T000000Z.png
var img169 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T010000Z.png
var img170 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T020000Z.png
var img171 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T030000Z.png
var img172 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T040000Z.png
var img173 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T050000Z.png
var img174 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T060000Z.png
var img175 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T070000Z.png
var img176 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T080000Z.png
var img177 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T090000Z.png
var img178 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T100000Z.png
var img179 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T110000Z.png
var img180 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T120000Z.png
var img181 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T130000Z.png
var img182 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T140000Z.png
var img183 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T150000Z.png
var img184 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T160000Z.png
var img185 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T170000Z.png
var img186 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T180000Z.png
var img187 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T190000Z.png
var img188 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T200000Z.png
var img189 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T210000Z.png
var img190 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T220000Z.png
var img191 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950108T230000Z.png
var img192 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T000000Z.png
var img193 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T010000Z.png
var img194 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T020000Z.png
var img195 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T030000Z.png
var img196 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T040000Z.png
var img197 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T050000Z.png
var img198 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T060000Z.png
var img199 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T070000Z.png
var img200 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T080000Z.png
var img201 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T090000Z.png
var img202 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T100000Z.png
var img203 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T110000Z.png
var img204 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T120000Z.png
var img205 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T130000Z.png
var img206 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T140000Z.png
var img207 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T150000Z.png
var img208 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T160000Z.png
var img209 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T170000Z.png
var img210 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T180000Z.png
var img211 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T190000Z.png
var img212 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T200000Z.png
var img213 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T210000Z.png
var img214 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T220000Z.png
var img215 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950109T230000Z.png
var img216 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T000000Z.png
var img217 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T010000Z.png
var img218 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T020000Z.png
var img219 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T030000Z.png
var img220 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T040000Z.png
var img221 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T050000Z.png
var img222 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T060000Z.png
var img223 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T070000Z.png
var img224 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T080000Z.png
var img225 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T090000Z.png
var img226 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T100000Z.png
var img227 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T110000Z.png
var img228 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T120000Z.png
var img229 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T130000Z.png
var img230 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T140000Z.png
var img231 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T150000Z.png
var img232 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T160000Z.png
var img233 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T170000Z.png
var img234 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T180000Z.png
var img235 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T190000Z.png
var img236 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T200000Z.png
var img237 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T210000Z.png
var img238 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T220000Z.png
var img239 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950110T230000Z.png
var img240 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T000000Z.png
var img241 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T010000Z.png
var img242 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T020000Z.png
var img243 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T030000Z.png
var img244 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T040000Z.png
var img245 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T050000Z.png
var img246 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T060000Z.png
var img247 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T070000Z.png
var img248 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T080000Z.png
var img249 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T090000Z.png
var img250 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T100000Z.png
var img251 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T110000Z.png
var img252 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T120000Z.png
var img253 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T130000Z.png
var img254 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T140000Z.png
var img255 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T150000Z.png
var img256 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T160000Z.png
var img257 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T170000Z.png
var img258 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T180000Z.png
var img259 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T190000Z.png
var img260 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T200000Z.png
var img261 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T210000Z.png
var img262 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T220000Z.png
var img263 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950111T230000Z.png
var img264 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T000000Z.png
var img265 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T010000Z.png
var img266 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T020000Z.png
var img267 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T030000Z.png
var img268 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T040000Z.png
var img269 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T050000Z.png
var img270 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T060000Z.png
var img271 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T070000Z.png
var img272 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T080000Z.png
var img273 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T090000Z.png
var img274 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T100000Z.png
var img275 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T110000Z.png
var img276 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T120000Z.png
var img277 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T130000Z.png
var img278 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T140000Z.png
var img279 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T150000Z.png
var img280 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T160000Z.png
var img281 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T170000Z.png
var img282 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T180000Z.png
var img283 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T190000Z.png
var img284 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T200000Z.png
var img285 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T210000Z.png
var img286 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T220000Z.png
var img287 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950112T230000Z.png
var img288 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T000000Z.png
var img289 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T010000Z.png
var img290 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T020000Z.png
var img291 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T030000Z.png
var img292 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T040000Z.png
var img293 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T050000Z.png
var img294 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T060000Z.png
var img295 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T070000Z.png
var img296 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T080000Z.png
var img297 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T090000Z.png
var img298 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T100000Z.png
var img299 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T110000Z.png
var img300 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T120000Z.png
var img301 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T130000Z.png
var img302 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T140000Z.png
var img303 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T150000Z.png
var img304 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T160000Z.png
var img305 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T170000Z.png
var img306 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T180000Z.png
var img307 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T190000Z.png
var img308 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T200000Z.png
var img309 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T210000Z.png
var img310 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T220000Z.png
var img311 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950113T230000Z.png
var img312 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T000000Z.png
var img313 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T010000Z.png
var img314 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T020000Z.png
var img315 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T030000Z.png
var img316 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T040000Z.png
var img317 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T050000Z.png
var img318 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T060000Z.png
var img319 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T070000Z.png
var img320 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T080000Z.png
var img321 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T090000Z.png
var img322 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T100000Z.png
var img323 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T110000Z.png
var img324 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T120000Z.png
var img325 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T130000Z.png
var img326 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T140000Z.png
var img327 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T150000Z.png
var img328 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T160000Z.png
var img329 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T170000Z.png
var img330 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T180000Z.png
var img331 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T190000Z.png
var img332 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T200000Z.png
var img333 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T210000Z.png
var img334 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T220000Z.png
var img335 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950114T230000Z.png
var img336 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T000000Z.png
var img337 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T010000Z.png
var img338 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T020000Z.png
var img339 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T030000Z.png
var img340 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T040000Z.png
var img341 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T050000Z.png
var img342 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T060000Z.png
var img343 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T070000Z.png
var img344 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T080000Z.png
var img345 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T090000Z.png
var img346 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T100000Z.png
var img347 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T110000Z.png
var img348 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T120000Z.png
var img349 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T130000Z.png
var img350 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T140000Z.png
var img351 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T150000Z.png
var img352 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T160000Z.png
var img353 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T170000Z.png
var img354 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T180000Z.png
var img355 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T190000Z.png
var img356 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T200000Z.png
var img357 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T210000Z.png
var img358 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T220000Z.png
var img359 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950115T230000Z.png
var img360 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T000000Z.png
var img361 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T010000Z.png
var img362 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T020000Z.png
var img363 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T030000Z.png
var img364 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T040000Z.png
var img365 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T050000Z.png
var img366 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T060000Z.png
var img367 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T070000Z.png
var img368 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T080000Z.png
var img369 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T090000Z.png
var img370 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T100000Z.png
var img371 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T110000Z.png
var img372 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T120000Z.png
var img373 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T130000Z.png
var img374 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T140000Z.png
var img375 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T150000Z.png
var img376 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T160000Z.png
var img377 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T170000Z.png
var img378 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T180000Z.png
var img379 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T190000Z.png
var img380 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T200000Z.png
var img381 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T210000Z.png
var img382 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T220000Z.png
var img383 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950116T230000Z.png
var img384 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T000000Z.png
var img385 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T010000Z.png
var img386 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T020000Z.png
var img387 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T030000Z.png
var img388 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T040000Z.png
var img389 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T050000Z.png
var img390 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T060000Z.png
var img391 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T070000Z.png
var img392 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T080000Z.png
var img393 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T090000Z.png
var img394 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T100000Z.png
var img395 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T110000Z.png
var img396 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T120000Z.png
var img397 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T130000Z.png
var img398 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T140000Z.png
var img399 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T150000Z.png
var img400 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T160000Z.png
var img401 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T170000Z.png
var img402 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T180000Z.png
var img403 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T190000Z.png
var img404 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T200000Z.png
var img405 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T210000Z.png
var img406 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T220000Z.png
var img407 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950117T230000Z.png
var img408 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T000000Z.png
var img409 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T010000Z.png
var img410 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T020000Z.png
var img411 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T030000Z.png
var img412 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T040000Z.png
var img413 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T050000Z.png
var img414 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T060000Z.png
var img415 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T070000Z.png
var img416 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T080000Z.png
var img417 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T090000Z.png
var img418 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T100000Z.png
var img419 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T110000Z.png
var img420 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T120000Z.png
var img421 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T130000Z.png
var img422 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T140000Z.png
var img423 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T150000Z.png
var img424 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T160000Z.png
var img425 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T170000Z.png
var img426 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T180000Z.png
var img427 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T190000Z.png
var img428 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T200000Z.png
var img429 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T210000Z.png
var img430 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T220000Z.png
var img431 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950118T230000Z.png
var img432 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T000000Z.png
var img433 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T010000Z.png
var img434 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T020000Z.png
var img435 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T030000Z.png
var img436 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T040000Z.png
var img437 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T050000Z.png
var img438 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T060000Z.png
var img439 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T070000Z.png
var img440 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T080000Z.png
var img441 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T090000Z.png
var img442 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T100000Z.png
var img443 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T110000Z.png
var img444 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T120000Z.png
var img445 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T130000Z.png
var img446 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T140000Z.png
var img447 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T150000Z.png
var img448 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T160000Z.png
var img449 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T170000Z.png
var img450 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T180000Z.png
var img451 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T190000Z.png
var img452 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T200000Z.png
var img453 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T210000Z.png
var img454 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T220000Z.png
var img455 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950119T230000Z.png
var img456 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T000000Z.png
var img457 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T010000Z.png
var img458 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T020000Z.png
var img459 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T030000Z.png
var img460 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T040000Z.png
var img461 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T050000Z.png
var img462 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T060000Z.png
var img463 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T070000Z.png
var img464 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T080000Z.png
var img465 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T090000Z.png
var img466 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T100000Z.png
var img467 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T110000Z.png
var img468 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T120000Z.png
var img469 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T130000Z.png
var img470 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T140000Z.png
var img471 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T150000Z.png
var img472 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T160000Z.png
var img473 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T170000Z.png
var img474 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T180000Z.png
var img475 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T190000Z.png
var img476 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T200000Z.png
var img477 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T210000Z.png
var img478 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T220000Z.png
var img479 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950120T230000Z.png
var img480 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T000000Z.png
var img481 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T010000Z.png
var img482 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T020000Z.png
var img483 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T030000Z.png
var img484 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T040000Z.png
var img485 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T050000Z.png
var img486 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T060000Z.png
var img487 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T070000Z.png
var img488 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T080000Z.png
var img489 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T090000Z.png
var img490 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T100000Z.png
var img491 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T110000Z.png
var img492 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T120000Z.png
var img493 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T130000Z.png
var img494 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T140000Z.png
var img495 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T150000Z.png
var img496 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T160000Z.png
var img497 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T170000Z.png
var img498 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T180000Z.png
var img499 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T190000Z.png
var img500 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T200000Z.png
var img501 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T210000Z.png
var img502 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T220000Z.png
var img503 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950121T230000Z.png
var img504 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T000000Z.png
var img505 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T010000Z.png
var img506 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T020000Z.png
var img507 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T030000Z.png
var img508 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T040000Z.png
var img509 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T050000Z.png
var img510 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T060000Z.png
var img511 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T070000Z.png
var img512 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T080000Z.png
var img513 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T090000Z.png
var img514 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T100000Z.png
var img515 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T110000Z.png
var img516 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T120000Z.png
var img517 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T130000Z.png
var img518 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T140000Z.png
var img519 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T150000Z.png
var img520 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T160000Z.png
var img521 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T170000Z.png
var img522 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T180000Z.png
var img523 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T190000Z.png
var img524 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T200000Z.png
var img525 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T210000Z.png
var img526 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T220000Z.png
var img527 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950122T230000Z.png
var img528 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T000000Z.png
var img529 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T010000Z.png
var img530 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T020000Z.png
var img531 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T030000Z.png
var img532 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T040000Z.png
var img533 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T050000Z.png
var img534 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T060000Z.png
var img535 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T070000Z.png
var img536 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T080000Z.png
var img537 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T090000Z.png
var img538 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T100000Z.png
var img539 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T110000Z.png
var img540 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T120000Z.png
var img541 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T130000Z.png
var img542 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T140000Z.png
var img543 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T150000Z.png
var img544 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T160000Z.png
var img545 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T170000Z.png
var img546 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T180000Z.png
var img547 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T190000Z.png
var img548 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T200000Z.png
var img549 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T210000Z.png
var img550 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T220000Z.png
var img551 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950123T230000Z.png
var img552 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T000000Z.png
var img553 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T010000Z.png
var img554 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T020000Z.png
var img555 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T030000Z.png
var img556 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T040000Z.png
var img557 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T050000Z.png
var img558 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T060000Z.png
var img559 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T070000Z.png
var img560 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T080000Z.png
var img561 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T090000Z.png
var img562 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T100000Z.png
var img563 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T110000Z.png
var img564 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T120000Z.png
var img565 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T130000Z.png
var img566 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T140000Z.png
var img567 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T150000Z.png
var img568 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T160000Z.png
var img569 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T170000Z.png
var img570 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T180000Z.png
var img571 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T190000Z.png
var img572 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T200000Z.png
var img573 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T210000Z.png
var img574 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T220000Z.png
var img575 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950124T230000Z.png
var img576 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T000000Z.png
var img577 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T010000Z.png
var img578 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T020000Z.png
var img579 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T030000Z.png
var img580 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T040000Z.png
var img581 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T050000Z.png
var img582 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T060000Z.png
var img583 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T070000Z.png
var img584 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T080000Z.png
var img585 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T090000Z.png
var img586 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T100000Z.png
var img587 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T110000Z.png
var img588 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T120000Z.png
var img589 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T130000Z.png
var img590 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T140000Z.png
var img591 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T150000Z.png
var img592 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T160000Z.png
var img593 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T170000Z.png
var img594 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T180000Z.png
var img595 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T190000Z.png
var img596 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T200000Z.png
var img597 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T210000Z.png
var img598 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T220000Z.png
var img599 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950125T230000Z.png
var img600 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T000000Z.png
var img601 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T010000Z.png
var img602 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T020000Z.png
var img603 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T030000Z.png
var img604 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T040000Z.png
var img605 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T050000Z.png
var img606 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T060000Z.png
var img607 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T070000Z.png
var img608 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T080000Z.png
var img609 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T090000Z.png
var img610 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T100000Z.png
var img611 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T110000Z.png
var img612 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T120000Z.png
var img613 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T130000Z.png
var img614 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T140000Z.png
var img615 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T150000Z.png
var img616 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T160000Z.png
var img617 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T170000Z.png
var img618 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T180000Z.png
var img619 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T190000Z.png
var img620 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T200000Z.png
var img621 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T210000Z.png
var img622 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T220000Z.png
var img623 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950126T230000Z.png
var img624 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T000000Z.png
var img625 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T010000Z.png
var img626 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T020000Z.png
var img627 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T030000Z.png
var img628 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T040000Z.png
var img629 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T050000Z.png
var img630 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T060000Z.png
var img631 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T070000Z.png
var img632 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T080000Z.png
var img633 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T090000Z.png
var img634 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T100000Z.png
var img635 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T110000Z.png
var img636 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T120000Z.png
var img637 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T130000Z.png
var img638 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T140000Z.png
var img639 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T150000Z.png
var img640 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T160000Z.png
var img641 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T170000Z.png
var img642 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T180000Z.png
var img643 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T190000Z.png
var img644 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T200000Z.png
var img645 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T210000Z.png
var img646 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T220000Z.png
var img647 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950127T230000Z.png
var img648 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T000000Z.png
var img649 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T010000Z.png
var img650 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T020000Z.png
var img651 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T030000Z.png
var img652 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T040000Z.png
var img653 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T050000Z.png
var img654 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T060000Z.png
var img655 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T070000Z.png
var img656 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T080000Z.png
var img657 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T090000Z.png
var img658 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T100000Z.png
var img659 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T110000Z.png
var img660 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T120000Z.png
var img661 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T130000Z.png
var img662 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T140000Z.png
var img663 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T150000Z.png
var img664 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T160000Z.png
var img665 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T170000Z.png
var img666 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T180000Z.png
var img667 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T190000Z.png
var img668 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T200000Z.png
var img669 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T210000Z.png
var img670 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T220000Z.png
var img671 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950128T230000Z.png
var img672 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T000000Z.png
var img673 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T010000Z.png
var img674 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T020000Z.png
var img675 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T030000Z.png
var img676 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T040000Z.png
var img677 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T050000Z.png
var img678 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T060000Z.png
var img679 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T070000Z.png
var img680 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T080000Z.png
var img681 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T090000Z.png
var img682 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T100000Z.png
var img683 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T110000Z.png
var img684 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T120000Z.png
var img685 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T130000Z.png
var img686 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T140000Z.png
var img687 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T150000Z.png
var img688 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T160000Z.png
var img689 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T170000Z.png
var img690 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T180000Z.png
var img691 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T190000Z.png
var img692 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T200000Z.png
var img693 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T210000Z.png
var img694 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T220000Z.png
var img695 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950129T230000Z.png
var img696 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T000000Z.png
var img697 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T010000Z.png
var img698 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T020000Z.png
var img699 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T030000Z.png
var img700 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T040000Z.png
var img701 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T050000Z.png
var img702 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T060000Z.png
var img703 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T070000Z.png
var img704 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T080000Z.png
var img705 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T090000Z.png
var img706 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T100000Z.png
var img707 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T110000Z.png
var img708 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T120000Z.png
var img709 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T130000Z.png
var img710 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T140000Z.png
var img711 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T150000Z.png
var img712 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T160000Z.png
var img713 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T170000Z.png
var img714 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T180000Z.png
var img715 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T190000Z.png
var img716 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T200000Z.png
var img717 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T210000Z.png
var img718 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T220000Z.png
var img719 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950130T230000Z.png
var img720 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T000000Z.png
var img721 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T010000Z.png
var img722 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T020000Z.png
var img723 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T030000Z.png
var img724 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T040000Z.png
var img725 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T050000Z.png
var img726 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T060000Z.png
var img727 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T070000Z.png
var img728 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T080000Z.png
var img729 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T090000Z.png
var img730 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T100000Z.png
var img731 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T110000Z.png
var img732 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T120000Z.png
var img733 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T130000Z.png
var img734 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T140000Z.png
var img735 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T150000Z.png
var img736 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T160000Z.png
var img737 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T170000Z.png
var img738 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T180000Z.png
var img739 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T190000Z.png
var img740 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T200000Z.png
var img741 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T210000Z.png
var img742 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T220000Z.png
var img743 []byte
//go:embed images/gridrad_3d_v3_1_col_max_refl_19950131T230000Z.png
var img744 []byte

var images = map[string][]byte{
	"gridrad_3d_v3_1_col_max_refl_19950101T000000Z.png": img1,
	"gridrad_3d_v3_1_col_max_refl_19950101T010000Z.png": img2,
	"gridrad_3d_v3_1_col_max_refl_19950101T020000Z.png": img3,
	"gridrad_3d_v3_1_col_max_refl_19950101T030000Z.png": img4,
	"gridrad_3d_v3_1_col_max_refl_19950101T040000Z.png": img5,
	"gridrad_3d_v3_1_col_max_refl_19950101T050000Z.png": img6,
	"gridrad_3d_v3_1_col_max_refl_19950101T060000Z.png": img7,
	"gridrad_3d_v3_1_col_max_refl_19950101T070000Z.png": img8,
	"gridrad_3d_v3_1_col_max_refl_19950101T080000Z.png": img9,
	"gridrad_3d_v3_1_col_max_refl_19950101T090000Z.png": img10,
	"gridrad_3d_v3_1_col_max_refl_19950101T100000Z.png": img11,
	"gridrad_3d_v3_1_col_max_refl_19950101T110000Z.png": img12,
	"gridrad_3d_v3_1_col_max_refl_19950101T120000Z.png": img13,
	"gridrad_3d_v3_1_col_max_refl_19950101T130000Z.png": img14,
	"gridrad_3d_v3_1_col_max_refl_19950101T140000Z.png": img15,
	"gridrad_3d_v3_1_col_max_refl_19950101T150000Z.png": img16,
	"gridrad_3d_v3_1_col_max_refl_19950101T160000Z.png": img17,
	"gridrad_3d_v3_1_col_max_refl_19950101T170000Z.png": img18,
	"gridrad_3d_v3_1_col_max_refl_19950101T180000Z.png": img19,
	"gridrad_3d_v3_1_col_max_refl_19950101T190000Z.png": img20,
	"gridrad_3d_v3_1_col_max_refl_19950101T200000Z.png": img21,
	"gridrad_3d_v3_1_col_max_refl_19950101T210000Z.png": img22,
	"gridrad_3d_v3_1_col_max_refl_19950101T220000Z.png": img23,
	"gridrad_3d_v3_1_col_max_refl_19950101T230000Z.png": img24,
	"gridrad_3d_v3_1_col_max_refl_19950102T000000Z.png": img25,
	"gridrad_3d_v3_1_col_max_refl_19950102T010000Z.png": img26,
	"gridrad_3d_v3_1_col_max_refl_19950102T020000Z.png": img27,
	"gridrad_3d_v3_1_col_max_refl_19950102T030000Z.png": img28,
	"gridrad_3d_v3_1_col_max_refl_19950102T040000Z.png": img29,
	"gridrad_3d_v3_1_col_max_refl_19950102T050000Z.png": img30,
	"gridrad_3d_v3_1_col_max_refl_19950102T060000Z.png": img31,
	"gridrad_3d_v3_1_col_max_refl_19950102T070000Z.png": img32,
	"gridrad_3d_v3_1_col_max_refl_19950102T080000Z.png": img33,
	"gridrad_3d_v3_1_col_max_refl_19950102T090000Z.png": img34,
	"gridrad_3d_v3_1_col_max_refl_19950102T100000Z.png": img35,
	"gridrad_3d_v3_1_col_max_refl_19950102T110000Z.png": img36,
	"gridrad_3d_v3_1_col_max_refl_19950102T120000Z.png": img37,
	"gridrad_3d_v3_1_col_max_refl_19950102T130000Z.png": img38,
	"gridrad_3d_v3_1_col_max_refl_19950102T140000Z.png": img39,
	"gridrad_3d_v3_1_col_max_refl_19950102T150000Z.png": img40,
	"gridrad_3d_v3_1_col_max_refl_19950102T160000Z.png": img41,
	"gridrad_3d_v3_1_col_max_refl_19950102T170000Z.png": img42,
	"gridrad_3d_v3_1_col_max_refl_19950102T180000Z.png": img43,
	"gridrad_3d_v3_1_col_max_refl_19950102T190000Z.png": img44,
	"gridrad_3d_v3_1_col_max_refl_19950102T200000Z.png": img45,
	"gridrad_3d_v3_1_col_max_refl_19950102T210000Z.png": img46,
	"gridrad_3d_v3_1_col_max_refl_19950102T220000Z.png": img47,
	"gridrad_3d_v3_1_col_max_refl_19950102T230000Z.png": img48,
	"gridrad_3d_v3_1_col_max_refl_19950103T000000Z.png": img49,
	"gridrad_3d_v3_1_col_max_refl_19950103T010000Z.png": img50,
	"gridrad_3d_v3_1_col_max_refl_19950103T020000Z.png": img51,
	"gridrad_3d_v3_1_col_max_refl_19950103T030000Z.png": img52,
	"gridrad_3d_v3_1_col_max_refl_19950103T040000Z.png": img53,
	"gridrad_3d_v3_1_col_max_refl_19950103T050000Z.png": img54,
	"gridrad_3d_v3_1_col_max_refl_19950103T060000Z.png": img55,
	"gridrad_3d_v3_1_col_max_refl_19950103T070000Z.png": img56,
	"gridrad_3d_v3_1_col_max_refl_19950103T080000Z.png": img57,
	"gridrad_3d_v3_1_col_max_refl_19950103T090000Z.png": img58,
	"gridrad_3d_v3_1_col_max_refl_19950103T100000Z.png": img59,
	"gridrad_3d_v3_1_col_max_refl_19950103T110000Z.png": img60,
	"gridrad_3d_v3_1_col_max_refl_19950103T120000Z.png": img61,
	"gridrad_3d_v3_1_col_max_refl_19950103T130000Z.png": img62,
	"gridrad_3d_v3_1_col_max_refl_19950103T140000Z.png": img63,
	"gridrad_3d_v3_1_col_max_refl_19950103T150000Z.png": img64,
	"gridrad_3d_v3_1_col_max_refl_19950103T160000Z.png": img65,
	"gridrad_3d_v3_1_col_max_refl_19950103T170000Z.png": img66,
	"gridrad_3d_v3_1_col_max_refl_19950103T180000Z.png": img67,
	"gridrad_3d_v3_1_col_max_refl_19950103T190000Z.png": img68,
	"gridrad_3d_v3_1_col_max_refl_19950103T200000Z.png": img69,
	"gridrad_3d_v3_1_col_max_refl_19950103T210000Z.png": img70,
	"gridrad_3d_v3_1_col_max_refl_19950103T220000Z.png": img71,
	"gridrad_3d_v3_1_col_max_refl_19950103T230000Z.png": img72,
	"gridrad_3d_v3_1_col_max_refl_19950104T000000Z.png": img73,
	"gridrad_3d_v3_1_col_max_refl_19950104T010000Z.png": img74,
	"gridrad_3d_v3_1_col_max_refl_19950104T020000Z.png": img75,
	"gridrad_3d_v3_1_col_max_refl_19950104T030000Z.png": img76,
	"gridrad_3d_v3_1_col_max_refl_19950104T040000Z.png": img77,
	"gridrad_3d_v3_1_col_max_refl_19950104T050000Z.png": img78,
	"gridrad_3d_v3_1_col_max_refl_19950104T060000Z.png": img79,
	"gridrad_3d_v3_1_col_max_refl_19950104T070000Z.png": img80,
	"gridrad_3d_v3_1_col_max_refl_19950104T080000Z.png": img81,
	"gridrad_3d_v3_1_col_max_refl_19950104T090000Z.png": img82,
	"gridrad_3d_v3_1_col_max_refl_19950104T100000Z.png": img83,
	"gridrad_3d_v3_1_col_max_refl_19950104T110000Z.png": img84,
	"gridrad_3d_v3_1_col_max_refl_19950104T120000Z.png": img85,
	"gridrad_3d_v3_1_col_max_refl_19950104T130000Z.png": img86,
	"gridrad_3d_v3_1_col_max_refl_19950104T140000Z.png": img87,
	"gridrad_3d_v3_1_col_max_refl_19950104T150000Z.png": img88,
	"gridrad_3d_v3_1_col_max_refl_19950104T160000Z.png": img89,
	"gridrad_3d_v3_1_col_max_refl_19950104T170000Z.png": img90,
	"gridrad_3d_v3_1_col_max_refl_19950104T180000Z.png": img91,
	"gridrad_3d_v3_1_col_max_refl_19950104T190000Z.png": img92,
	"gridrad_3d_v3_1_col_max_refl_19950104T200000Z.png": img93,
	"gridrad_3d_v3_1_col_max_refl_19950104T210000Z.png": img94,
	"gridrad_3d_v3_1_col_max_refl_19950104T220000Z.png": img95,
	"gridrad_3d_v3_1_col_max_refl_19950104T230000Z.png": img96,
	"gridrad_3d_v3_1_col_max_refl_19950105T000000Z.png": img97,
	"gridrad_3d_v3_1_col_max_refl_19950105T010000Z.png": img98,
	"gridrad_3d_v3_1_col_max_refl_19950105T020000Z.png": img99,
	"gridrad_3d_v3_1_col_max_refl_19950105T030000Z.png": img100,
	"gridrad_3d_v3_1_col_max_refl_19950105T040000Z.png": img101,
	"gridrad_3d_v3_1_col_max_refl_19950105T050000Z.png": img102,
	"gridrad_3d_v3_1_col_max_refl_19950105T060000Z.png": img103,
	"gridrad_3d_v3_1_col_max_refl_19950105T070000Z.png": img104,
	"gridrad_3d_v3_1_col_max_refl_19950105T080000Z.png": img105,
	"gridrad_3d_v3_1_col_max_refl_19950105T090000Z.png": img106,
	"gridrad_3d_v3_1_col_max_refl_19950105T100000Z.png": img107,
	"gridrad_3d_v3_1_col_max_refl_19950105T110000Z.png": img108,
	"gridrad_3d_v3_1_col_max_refl_19950105T120000Z.png": img109,
	"gridrad_3d_v3_1_col_max_refl_19950105T130000Z.png": img110,
	"gridrad_3d_v3_1_col_max_refl_19950105T140000Z.png": img111,
	"gridrad_3d_v3_1_col_max_refl_19950105T150000Z.png": img112,
	"gridrad_3d_v3_1_col_max_refl_19950105T160000Z.png": img113,
	"gridrad_3d_v3_1_col_max_refl_19950105T170000Z.png": img114,
	"gridrad_3d_v3_1_col_max_refl_19950105T180000Z.png": img115,
	"gridrad_3d_v3_1_col_max_refl_19950105T190000Z.png": img116,
	"gridrad_3d_v3_1_col_max_refl_19950105T200000Z.png": img117,
	"gridrad_3d_v3_1_col_max_refl_19950105T210000Z.png": img118,
	"gridrad_3d_v3_1_col_max_refl_19950105T220000Z.png": img119,
	"gridrad_3d_v3_1_col_max_refl_19950105T230000Z.png": img120,
	"gridrad_3d_v3_1_col_max_refl_19950106T000000Z.png": img121,
	"gridrad_3d_v3_1_col_max_refl_19950106T010000Z.png": img122,
	"gridrad_3d_v3_1_col_max_refl_19950106T020000Z.png": img123,
	"gridrad_3d_v3_1_col_max_refl_19950106T030000Z.png": img124,
	"gridrad_3d_v3_1_col_max_refl_19950106T040000Z.png": img125,
	"gridrad_3d_v3_1_col_max_refl_19950106T050000Z.png": img126,
	"gridrad_3d_v3_1_col_max_refl_19950106T060000Z.png": img127,
	"gridrad_3d_v3_1_col_max_refl_19950106T070000Z.png": img128,
	"gridrad_3d_v3_1_col_max_refl_19950106T080000Z.png": img129,
	"gridrad_3d_v3_1_col_max_refl_19950106T090000Z.png": img130,
	"gridrad_3d_v3_1_col_max_refl_19950106T100000Z.png": img131,
	"gridrad_3d_v3_1_col_max_refl_19950106T110000Z.png": img132,
	"gridrad_3d_v3_1_col_max_refl_19950106T120000Z.png": img133,
	"gridrad_3d_v3_1_col_max_refl_19950106T130000Z.png": img134,
	"gridrad_3d_v3_1_col_max_refl_19950106T140000Z.png": img135,
	"gridrad_3d_v3_1_col_max_refl_19950106T150000Z.png": img136,
	"gridrad_3d_v3_1_col_max_refl_19950106T160000Z.png": img137,
	"gridrad_3d_v3_1_col_max_refl_19950106T170000Z.png": img138,
	"gridrad_3d_v3_1_col_max_refl_19950106T180000Z.png": img139,
	"gridrad_3d_v3_1_col_max_refl_19950106T190000Z.png": img140,
	"gridrad_3d_v3_1_col_max_refl_19950106T200000Z.png": img141,
	"gridrad_3d_v3_1_col_max_refl_19950106T210000Z.png": img142,
	"gridrad_3d_v3_1_col_max_refl_19950106T220000Z.png": img143,
	"gridrad_3d_v3_1_col_max_refl_19950106T230000Z.png": img144,
	"gridrad_3d_v3_1_col_max_refl_19950107T000000Z.png": img145,
	"gridrad_3d_v3_1_col_max_refl_19950107T010000Z.png": img146,
	"gridrad_3d_v3_1_col_max_refl_19950107T020000Z.png": img147,
	"gridrad_3d_v3_1_col_max_refl_19950107T030000Z.png": img148,
	"gridrad_3d_v3_1_col_max_refl_19950107T040000Z.png": img149,
	"gridrad_3d_v3_1_col_max_refl_19950107T050000Z.png": img150,
	"gridrad_3d_v3_1_col_max_refl_19950107T060000Z.png": img151,
	"gridrad_3d_v3_1_col_max_refl_19950107T070000Z.png": img152,
	"gridrad_3d_v3_1_col_max_refl_19950107T080000Z.png": img153,
	"gridrad_3d_v3_1_col_max_refl_19950107T090000Z.png": img154,
	"gridrad_3d_v3_1_col_max_refl_19950107T100000Z.png": img155,
	"gridrad_3d_v3_1_col_max_refl_19950107T110000Z.png": img156,
	"gridrad_3d_v3_1_col_max_refl_19950107T120000Z.png": img157,
	"gridrad_3d_v3_1_col_max_refl_19950107T130000Z.png": img158,
	"gridrad_3d_v3_1_col_max_refl_19950107T140000Z.png": img159,
	"gridrad_3d_v3_1_col_max_refl_19950107T150000Z.png": img160,
	"gridrad_3d_v3_1_col_max_refl_19950107T160000Z.png": img161,
	"gridrad_3d_v3_1_col_max_refl_19950107T170000Z.png": img162,
	"gridrad_3d_v3_1_col_max_refl_19950107T180000Z.png": img163,
	"gridrad_3d_v3_1_col_max_refl_19950107T190000Z.png": img164,
	"gridrad_3d_v3_1_col_max_refl_19950107T200000Z.png": img165,
	"gridrad_3d_v3_1_col_max_refl_19950107T210000Z.png": img166,
	"gridrad_3d_v3_1_col_max_refl_19950107T220000Z.png": img167,
	"gridrad_3d_v3_1_col_max_refl_19950107T230000Z.png": img168,
	"gridrad_3d_v3_1_col_max_refl_19950108T000000Z.png": img169,
	"gridrad_3d_v3_1_col_max_refl_19950108T010000Z.png": img170,
	"gridrad_3d_v3_1_col_max_refl_19950108T020000Z.png": img171,
	"gridrad_3d_v3_1_col_max_refl_19950108T030000Z.png": img172,
	"gridrad_3d_v3_1_col_max_refl_19950108T040000Z.png": img173,
	"gridrad_3d_v3_1_col_max_refl_19950108T050000Z.png": img174,
	"gridrad_3d_v3_1_col_max_refl_19950108T060000Z.png": img175,
	"gridrad_3d_v3_1_col_max_refl_19950108T070000Z.png": img176,
	"gridrad_3d_v3_1_col_max_refl_19950108T080000Z.png": img177,
	"gridrad_3d_v3_1_col_max_refl_19950108T090000Z.png": img178,
	"gridrad_3d_v3_1_col_max_refl_19950108T100000Z.png": img179,
	"gridrad_3d_v3_1_col_max_refl_19950108T110000Z.png": img180,
	"gridrad_3d_v3_1_col_max_refl_19950108T120000Z.png": img181,
	"gridrad_3d_v3_1_col_max_refl_19950108T130000Z.png": img182,
	"gridrad_3d_v3_1_col_max_refl_19950108T140000Z.png": img183,
	"gridrad_3d_v3_1_col_max_refl_19950108T150000Z.png": img184,
	"gridrad_3d_v3_1_col_max_refl_19950108T160000Z.png": img185,
	"gridrad_3d_v3_1_col_max_refl_19950108T170000Z.png": img186,
	"gridrad_3d_v3_1_col_max_refl_19950108T180000Z.png": img187,
	"gridrad_3d_v3_1_col_max_refl_19950108T190000Z.png": img188,
	"gridrad_3d_v3_1_col_max_refl_19950108T200000Z.png": img189,
	"gridrad_3d_v3_1_col_max_refl_19950108T210000Z.png": img190,
	"gridrad_3d_v3_1_col_max_refl_19950108T220000Z.png": img191,
	"gridrad_3d_v3_1_col_max_refl_19950108T230000Z.png": img192,
	"gridrad_3d_v3_1_col_max_refl_19950109T000000Z.png": img193,
	"gridrad_3d_v3_1_col_max_refl_19950109T010000Z.png": img194,
	"gridrad_3d_v3_1_col_max_refl_19950109T020000Z.png": img195,
	"gridrad_3d_v3_1_col_max_refl_19950109T030000Z.png": img196,
	"gridrad_3d_v3_1_col_max_refl_19950109T040000Z.png": img197,
	"gridrad_3d_v3_1_col_max_refl_19950109T050000Z.png": img198,
	"gridrad_3d_v3_1_col_max_refl_19950109T060000Z.png": img199,
	"gridrad_3d_v3_1_col_max_refl_19950109T070000Z.png": img200,
	"gridrad_3d_v3_1_col_max_refl_19950109T080000Z.png": img201,
	"gridrad_3d_v3_1_col_max_refl_19950109T090000Z.png": img202,
	"gridrad_3d_v3_1_col_max_refl_19950109T100000Z.png": img203,
	"gridrad_3d_v3_1_col_max_refl_19950109T110000Z.png": img204,
	"gridrad_3d_v3_1_col_max_refl_19950109T120000Z.png": img205,
	"gridrad_3d_v3_1_col_max_refl_19950109T130000Z.png": img206,
	"gridrad_3d_v3_1_col_max_refl_19950109T140000Z.png": img207,
	"gridrad_3d_v3_1_col_max_refl_19950109T150000Z.png": img208,
	"gridrad_3d_v3_1_col_max_refl_19950109T160000Z.png": img209,
	"gridrad_3d_v3_1_col_max_refl_19950109T170000Z.png": img210,
	"gridrad_3d_v3_1_col_max_refl_19950109T180000Z.png": img211,
	"gridrad_3d_v3_1_col_max_refl_19950109T190000Z.png": img212,
	"gridrad_3d_v3_1_col_max_refl_19950109T200000Z.png": img213,
	"gridrad_3d_v3_1_col_max_refl_19950109T210000Z.png": img214,
	"gridrad_3d_v3_1_col_max_refl_19950109T220000Z.png": img215,
	"gridrad_3d_v3_1_col_max_refl_19950109T230000Z.png": img216,
	"gridrad_3d_v3_1_col_max_refl_19950110T000000Z.png": img217,
	"gridrad_3d_v3_1_col_max_refl_19950110T010000Z.png": img218,
	"gridrad_3d_v3_1_col_max_refl_19950110T020000Z.png": img219,
	"gridrad_3d_v3_1_col_max_refl_19950110T030000Z.png": img220,
	"gridrad_3d_v3_1_col_max_refl_19950110T040000Z.png": img221,
	"gridrad_3d_v3_1_col_max_refl_19950110T050000Z.png": img222,
	"gridrad_3d_v3_1_col_max_refl_19950110T060000Z.png": img223,
	"gridrad_3d_v3_1_col_max_refl_19950110T070000Z.png": img224,
	"gridrad_3d_v3_1_col_max_refl_19950110T080000Z.png": img225,
	"gridrad_3d_v3_1_col_max_refl_19950110T090000Z.png": img226,
	"gridrad_3d_v3_1_col_max_refl_19950110T100000Z.png": img227,
	"gridrad_3d_v3_1_col_max_refl_19950110T110000Z.png": img228,
	"gridrad_3d_v3_1_col_max_refl_19950110T120000Z.png": img229,
	"gridrad_3d_v3_1_col_max_refl_19950110T130000Z.png": img230,
	"gridrad_3d_v3_1_col_max_refl_19950110T140000Z.png": img231,
	"gridrad_3d_v3_1_col_max_refl_19950110T150000Z.png": img232,
	"gridrad_3d_v3_1_col_max_refl_19950110T160000Z.png": img233,
	"gridrad_3d_v3_1_col_max_refl_19950110T170000Z.png": img234,
	"gridrad_3d_v3_1_col_max_refl_19950110T180000Z.png": img235,
	"gridrad_3d_v3_1_col_max_refl_19950110T190000Z.png": img236,
	"gridrad_3d_v3_1_col_max_refl_19950110T200000Z.png": img237,
	"gridrad_3d_v3_1_col_max_refl_19950110T210000Z.png": img238,
	"gridrad_3d_v3_1_col_max_refl_19950110T220000Z.png": img239,
	"gridrad_3d_v3_1_col_max_refl_19950110T230000Z.png": img240,
	"gridrad_3d_v3_1_col_max_refl_19950111T000000Z.png": img241,
	"gridrad_3d_v3_1_col_max_refl_19950111T010000Z.png": img242,
	"gridrad_3d_v3_1_col_max_refl_19950111T020000Z.png": img243,
	"gridrad_3d_v3_1_col_max_refl_19950111T030000Z.png": img244,
	"gridrad_3d_v3_1_col_max_refl_19950111T040000Z.png": img245,
	"gridrad_3d_v3_1_col_max_refl_19950111T050000Z.png": img246,
	"gridrad_3d_v3_1_col_max_refl_19950111T060000Z.png": img247,
	"gridrad_3d_v3_1_col_max_refl_19950111T070000Z.png": img248,
	"gridrad_3d_v3_1_col_max_refl_19950111T080000Z.png": img249,
	"gridrad_3d_v3_1_col_max_refl_19950111T090000Z.png": img250,
	"gridrad_3d_v3_1_col_max_refl_19950111T100000Z.png": img251,
	"gridrad_3d_v3_1_col_max_refl_19950111T110000Z.png": img252,
	"gridrad_3d_v3_1_col_max_refl_19950111T120000Z.png": img253,
	"gridrad_3d_v3_1_col_max_refl_19950111T130000Z.png": img254,
	"gridrad_3d_v3_1_col_max_refl_19950111T140000Z.png": img255,
	"gridrad_3d_v3_1_col_max_refl_19950111T150000Z.png": img256,
	"gridrad_3d_v3_1_col_max_refl_19950111T160000Z.png": img257,
	"gridrad_3d_v3_1_col_max_refl_19950111T170000Z.png": img258,
	"gridrad_3d_v3_1_col_max_refl_19950111T180000Z.png": img259,
	"gridrad_3d_v3_1_col_max_refl_19950111T190000Z.png": img260,
	"gridrad_3d_v3_1_col_max_refl_19950111T200000Z.png": img261,
	"gridrad_3d_v3_1_col_max_refl_19950111T210000Z.png": img262,
	"gridrad_3d_v3_1_col_max_refl_19950111T220000Z.png": img263,
	"gridrad_3d_v3_1_col_max_refl_19950111T230000Z.png": img264,
	"gridrad_3d_v3_1_col_max_refl_19950112T000000Z.png": img265,
	"gridrad_3d_v3_1_col_max_refl_19950112T010000Z.png": img266,
	"gridrad_3d_v3_1_col_max_refl_19950112T020000Z.png": img267,
	"gridrad_3d_v3_1_col_max_refl_19950112T030000Z.png": img268,
	"gridrad_3d_v3_1_col_max_refl_19950112T040000Z.png": img269,
	"gridrad_3d_v3_1_col_max_refl_19950112T050000Z.png": img270,
	"gridrad_3d_v3_1_col_max_refl_19950112T060000Z.png": img271,
	"gridrad_3d_v3_1_col_max_refl_19950112T070000Z.png": img272,
	"gridrad_3d_v3_1_col_max_refl_19950112T080000Z.png": img273,
	"gridrad_3d_v3_1_col_max_refl_19950112T090000Z.png": img274,
	"gridrad_3d_v3_1_col_max_refl_19950112T100000Z.png": img275,
	"gridrad_3d_v3_1_col_max_refl_19950112T110000Z.png": img276,
	"gridrad_3d_v3_1_col_max_refl_19950112T120000Z.png": img277,
	"gridrad_3d_v3_1_col_max_refl_19950112T130000Z.png": img278,
	"gridrad_3d_v3_1_col_max_refl_19950112T140000Z.png": img279,
	"gridrad_3d_v3_1_col_max_refl_19950112T150000Z.png": img280,
	"gridrad_3d_v3_1_col_max_refl_19950112T160000Z.png": img281,
	"gridrad_3d_v3_1_col_max_refl_19950112T170000Z.png": img282,
	"gridrad_3d_v3_1_col_max_refl_19950112T180000Z.png": img283,
	"gridrad_3d_v3_1_col_max_refl_19950112T190000Z.png": img284,
	"gridrad_3d_v3_1_col_max_refl_19950112T200000Z.png": img285,
	"gridrad_3d_v3_1_col_max_refl_19950112T210000Z.png": img286,
	"gridrad_3d_v3_1_col_max_refl_19950112T220000Z.png": img287,
	"gridrad_3d_v3_1_col_max_refl_19950112T230000Z.png": img288,
	"gridrad_3d_v3_1_col_max_refl_19950113T000000Z.png": img289,
	"gridrad_3d_v3_1_col_max_refl_19950113T010000Z.png": img290,
	"gridrad_3d_v3_1_col_max_refl_19950113T020000Z.png": img291,
	"gridrad_3d_v3_1_col_max_refl_19950113T030000Z.png": img292,
	"gridrad_3d_v3_1_col_max_refl_19950113T040000Z.png": img293,
	"gridrad_3d_v3_1_col_max_refl_19950113T050000Z.png": img294,
	"gridrad_3d_v3_1_col_max_refl_19950113T060000Z.png": img295,
	"gridrad_3d_v3_1_col_max_refl_19950113T070000Z.png": img296,
	"gridrad_3d_v3_1_col_max_refl_19950113T080000Z.png": img297,
	"gridrad_3d_v3_1_col_max_refl_19950113T090000Z.png": img298,
	"gridrad_3d_v3_1_col_max_refl_19950113T100000Z.png": img299,
	"gridrad_3d_v3_1_col_max_refl_19950113T110000Z.png": img300,
	"gridrad_3d_v3_1_col_max_refl_19950113T120000Z.png": img301,
	"gridrad_3d_v3_1_col_max_refl_19950113T130000Z.png": img302,
	"gridrad_3d_v3_1_col_max_refl_19950113T140000Z.png": img303,
	"gridrad_3d_v3_1_col_max_refl_19950113T150000Z.png": img304,
	"gridrad_3d_v3_1_col_max_refl_19950113T160000Z.png": img305,
	"gridrad_3d_v3_1_col_max_refl_19950113T170000Z.png": img306,
	"gridrad_3d_v3_1_col_max_refl_19950113T180000Z.png": img307,
	"gridrad_3d_v3_1_col_max_refl_19950113T190000Z.png": img308,
	"gridrad_3d_v3_1_col_max_refl_19950113T200000Z.png": img309,
	"gridrad_3d_v3_1_col_max_refl_19950113T210000Z.png": img310,
	"gridrad_3d_v3_1_col_max_refl_19950113T220000Z.png": img311,
	"gridrad_3d_v3_1_col_max_refl_19950113T230000Z.png": img312,
	"gridrad_3d_v3_1_col_max_refl_19950114T000000Z.png": img313,
	"gridrad_3d_v3_1_col_max_refl_19950114T010000Z.png": img314,
	"gridrad_3d_v3_1_col_max_refl_19950114T020000Z.png": img315,
	"gridrad_3d_v3_1_col_max_refl_19950114T030000Z.png": img316,
	"gridrad_3d_v3_1_col_max_refl_19950114T040000Z.png": img317,
	"gridrad_3d_v3_1_col_max_refl_19950114T050000Z.png": img318,
	"gridrad_3d_v3_1_col_max_refl_19950114T060000Z.png": img319,
	"gridrad_3d_v3_1_col_max_refl_19950114T070000Z.png": img320,
	"gridrad_3d_v3_1_col_max_refl_19950114T080000Z.png": img321,
	"gridrad_3d_v3_1_col_max_refl_19950114T090000Z.png": img322,
	"gridrad_3d_v3_1_col_max_refl_19950114T100000Z.png": img323,
	"gridrad_3d_v3_1_col_max_refl_19950114T110000Z.png": img324,
	"gridrad_3d_v3_1_col_max_refl_19950114T120000Z.png": img325,
	"gridrad_3d_v3_1_col_max_refl_19950114T130000Z.png": img326,
	"gridrad_3d_v3_1_col_max_refl_19950114T140000Z.png": img327,
	"gridrad_3d_v3_1_col_max_refl_19950114T150000Z.png": img328,
	"gridrad_3d_v3_1_col_max_refl_19950114T160000Z.png": img329,
	"gridrad_3d_v3_1_col_max_refl_19950114T170000Z.png": img330,
	"gridrad_3d_v3_1_col_max_refl_19950114T180000Z.png": img331,
	"gridrad_3d_v3_1_col_max_refl_19950114T190000Z.png": img332,
	"gridrad_3d_v3_1_col_max_refl_19950114T200000Z.png": img333,
	"gridrad_3d_v3_1_col_max_refl_19950114T210000Z.png": img334,
	"gridrad_3d_v3_1_col_max_refl_19950114T220000Z.png": img335,
	"gridrad_3d_v3_1_col_max_refl_19950114T230000Z.png": img336,
	"gridrad_3d_v3_1_col_max_refl_19950115T000000Z.png": img337,
	"gridrad_3d_v3_1_col_max_refl_19950115T010000Z.png": img338,
	"gridrad_3d_v3_1_col_max_refl_19950115T020000Z.png": img339,
	"gridrad_3d_v3_1_col_max_refl_19950115T030000Z.png": img340,
	"gridrad_3d_v3_1_col_max_refl_19950115T040000Z.png": img341,
	"gridrad_3d_v3_1_col_max_refl_19950115T050000Z.png": img342,
	"gridrad_3d_v3_1_col_max_refl_19950115T060000Z.png": img343,
	"gridrad_3d_v3_1_col_max_refl_19950115T070000Z.png": img344,
	"gridrad_3d_v3_1_col_max_refl_19950115T080000Z.png": img345,
	"gridrad_3d_v3_1_col_max_refl_19950115T090000Z.png": img346,
	"gridrad_3d_v3_1_col_max_refl_19950115T100000Z.png": img347,
	"gridrad_3d_v3_1_col_max_refl_19950115T110000Z.png": img348,
	"gridrad_3d_v3_1_col_max_refl_19950115T120000Z.png": img349,
	"gridrad_3d_v3_1_col_max_refl_19950115T130000Z.png": img350,
	"gridrad_3d_v3_1_col_max_refl_19950115T140000Z.png": img351,
	"gridrad_3d_v3_1_col_max_refl_19950115T150000Z.png": img352,
	"gridrad_3d_v3_1_col_max_refl_19950115T160000Z.png": img353,
	"gridrad_3d_v3_1_col_max_refl_19950115T170000Z.png": img354,
	"gridrad_3d_v3_1_col_max_refl_19950115T180000Z.png": img355,
	"gridrad_3d_v3_1_col_max_refl_19950115T190000Z.png": img356,
	"gridrad_3d_v3_1_col_max_refl_19950115T200000Z.png": img357,
	"gridrad_3d_v3_1_col_max_refl_19950115T210000Z.png": img358,
	"gridrad_3d_v3_1_col_max_refl_19950115T220000Z.png": img359,
	"gridrad_3d_v3_1_col_max_refl_19950115T230000Z.png": img360,
	"gridrad_3d_v3_1_col_max_refl_19950116T000000Z.png": img361,
	"gridrad_3d_v3_1_col_max_refl_19950116T010000Z.png": img362,
	"gridrad_3d_v3_1_col_max_refl_19950116T020000Z.png": img363,
	"gridrad_3d_v3_1_col_max_refl_19950116T030000Z.png": img364,
	"gridrad_3d_v3_1_col_max_refl_19950116T040000Z.png": img365,
	"gridrad_3d_v3_1_col_max_refl_19950116T050000Z.png": img366,
	"gridrad_3d_v3_1_col_max_refl_19950116T060000Z.png": img367,
	"gridrad_3d_v3_1_col_max_refl_19950116T070000Z.png": img368,
	"gridrad_3d_v3_1_col_max_refl_19950116T080000Z.png": img369,
	"gridrad_3d_v3_1_col_max_refl_19950116T090000Z.png": img370,
	"gridrad_3d_v3_1_col_max_refl_19950116T100000Z.png": img371,
	"gridrad_3d_v3_1_col_max_refl_19950116T110000Z.png": img372,
	"gridrad_3d_v3_1_col_max_refl_19950116T120000Z.png": img373,
	"gridrad_3d_v3_1_col_max_refl_19950116T130000Z.png": img374,
	"gridrad_3d_v3_1_col_max_refl_19950116T140000Z.png": img375,
	"gridrad_3d_v3_1_col_max_refl_19950116T150000Z.png": img376,
	"gridrad_3d_v3_1_col_max_refl_19950116T160000Z.png": img377,
	"gridrad_3d_v3_1_col_max_refl_19950116T170000Z.png": img378,
	"gridrad_3d_v3_1_col_max_refl_19950116T180000Z.png": img379,
	"gridrad_3d_v3_1_col_max_refl_19950116T190000Z.png": img380,
	"gridrad_3d_v3_1_col_max_refl_19950116T200000Z.png": img381,
	"gridrad_3d_v3_1_col_max_refl_19950116T210000Z.png": img382,
	"gridrad_3d_v3_1_col_max_refl_19950116T220000Z.png": img383,
	"gridrad_3d_v3_1_col_max_refl_19950116T230000Z.png": img384,
	"gridrad_3d_v3_1_col_max_refl_19950117T000000Z.png": img385,
	"gridrad_3d_v3_1_col_max_refl_19950117T010000Z.png": img386,
	"gridrad_3d_v3_1_col_max_refl_19950117T020000Z.png": img387,
	"gridrad_3d_v3_1_col_max_refl_19950117T030000Z.png": img388,
	"gridrad_3d_v3_1_col_max_refl_19950117T040000Z.png": img389,
	"gridrad_3d_v3_1_col_max_refl_19950117T050000Z.png": img390,
	"gridrad_3d_v3_1_col_max_refl_19950117T060000Z.png": img391,
	"gridrad_3d_v3_1_col_max_refl_19950117T070000Z.png": img392,
	"gridrad_3d_v3_1_col_max_refl_19950117T080000Z.png": img393,
	"gridrad_3d_v3_1_col_max_refl_19950117T090000Z.png": img394,
	"gridrad_3d_v3_1_col_max_refl_19950117T100000Z.png": img395,
	"gridrad_3d_v3_1_col_max_refl_19950117T110000Z.png": img396,
	"gridrad_3d_v3_1_col_max_refl_19950117T120000Z.png": img397,
	"gridrad_3d_v3_1_col_max_refl_19950117T130000Z.png": img398,
	"gridrad_3d_v3_1_col_max_refl_19950117T140000Z.png": img399,
	"gridrad_3d_v3_1_col_max_refl_19950117T150000Z.png": img400,
	"gridrad_3d_v3_1_col_max_refl_19950117T160000Z.png": img401,
	"gridrad_3d_v3_1_col_max_refl_19950117T170000Z.png": img402,
	"gridrad_3d_v3_1_col_max_refl_19950117T180000Z.png": img403,
	"gridrad_3d_v3_1_col_max_refl_19950117T190000Z.png": img404,
	"gridrad_3d_v3_1_col_max_refl_19950117T200000Z.png": img405,
	"gridrad_3d_v3_1_col_max_refl_19950117T210000Z.png": img406,
	"gridrad_3d_v3_1_col_max_refl_19950117T220000Z.png": img407,
	"gridrad_3d_v3_1_col_max_refl_19950117T230000Z.png": img408,
	"gridrad_3d_v3_1_col_max_refl_19950118T000000Z.png": img409,
	"gridrad_3d_v3_1_col_max_refl_19950118T010000Z.png": img410,
	"gridrad_3d_v3_1_col_max_refl_19950118T020000Z.png": img411,
	"gridrad_3d_v3_1_col_max_refl_19950118T030000Z.png": img412,
	"gridrad_3d_v3_1_col_max_refl_19950118T040000Z.png": img413,
	"gridrad_3d_v3_1_col_max_refl_19950118T050000Z.png": img414,
	"gridrad_3d_v3_1_col_max_refl_19950118T060000Z.png": img415,
	"gridrad_3d_v3_1_col_max_refl_19950118T070000Z.png": img416,
	"gridrad_3d_v3_1_col_max_refl_19950118T080000Z.png": img417,
	"gridrad_3d_v3_1_col_max_refl_19950118T090000Z.png": img418,
	"gridrad_3d_v3_1_col_max_refl_19950118T100000Z.png": img419,
	"gridrad_3d_v3_1_col_max_refl_19950118T110000Z.png": img420,
	"gridrad_3d_v3_1_col_max_refl_19950118T120000Z.png": img421,
	"gridrad_3d_v3_1_col_max_refl_19950118T130000Z.png": img422,
	"gridrad_3d_v3_1_col_max_refl_19950118T140000Z.png": img423,
	"gridrad_3d_v3_1_col_max_refl_19950118T150000Z.png": img424,
	"gridrad_3d_v3_1_col_max_refl_19950118T160000Z.png": img425,
	"gridrad_3d_v3_1_col_max_refl_19950118T170000Z.png": img426,
	"gridrad_3d_v3_1_col_max_refl_19950118T180000Z.png": img427,
	"gridrad_3d_v3_1_col_max_refl_19950118T190000Z.png": img428,
	"gridrad_3d_v3_1_col_max_refl_19950118T200000Z.png": img429,
	"gridrad_3d_v3_1_col_max_refl_19950118T210000Z.png": img430,
	"gridrad_3d_v3_1_col_max_refl_19950118T220000Z.png": img431,
	"gridrad_3d_v3_1_col_max_refl_19950118T230000Z.png": img432,
	"gridrad_3d_v3_1_col_max_refl_19950119T000000Z.png": img433,
	"gridrad_3d_v3_1_col_max_refl_19950119T010000Z.png": img434,
	"gridrad_3d_v3_1_col_max_refl_19950119T020000Z.png": img435,
	"gridrad_3d_v3_1_col_max_refl_19950119T030000Z.png": img436,
	"gridrad_3d_v3_1_col_max_refl_19950119T040000Z.png": img437,
	"gridrad_3d_v3_1_col_max_refl_19950119T050000Z.png": img438,
	"gridrad_3d_v3_1_col_max_refl_19950119T060000Z.png": img439,
	"gridrad_3d_v3_1_col_max_refl_19950119T070000Z.png": img440,
	"gridrad_3d_v3_1_col_max_refl_19950119T080000Z.png": img441,
	"gridrad_3d_v3_1_col_max_refl_19950119T090000Z.png": img442,
	"gridrad_3d_v3_1_col_max_refl_19950119T100000Z.png": img443,
	"gridrad_3d_v3_1_col_max_refl_19950119T110000Z.png": img444,
	"gridrad_3d_v3_1_col_max_refl_19950119T120000Z.png": img445,
	"gridrad_3d_v3_1_col_max_refl_19950119T130000Z.png": img446,
	"gridrad_3d_v3_1_col_max_refl_19950119T140000Z.png": img447,
	"gridrad_3d_v3_1_col_max_refl_19950119T150000Z.png": img448,
	"gridrad_3d_v3_1_col_max_refl_19950119T160000Z.png": img449,
	"gridrad_3d_v3_1_col_max_refl_19950119T170000Z.png": img450,
	"gridrad_3d_v3_1_col_max_refl_19950119T180000Z.png": img451,
	"gridrad_3d_v3_1_col_max_refl_19950119T190000Z.png": img452,
	"gridrad_3d_v3_1_col_max_refl_19950119T200000Z.png": img453,
	"gridrad_3d_v3_1_col_max_refl_19950119T210000Z.png": img454,
	"gridrad_3d_v3_1_col_max_refl_19950119T220000Z.png": img455,
	"gridrad_3d_v3_1_col_max_refl_19950119T230000Z.png": img456,
	"gridrad_3d_v3_1_col_max_refl_19950120T000000Z.png": img457,
	"gridrad_3d_v3_1_col_max_refl_19950120T010000Z.png": img458,
	"gridrad_3d_v3_1_col_max_refl_19950120T020000Z.png": img459,
	"gridrad_3d_v3_1_col_max_refl_19950120T030000Z.png": img460,
	"gridrad_3d_v3_1_col_max_refl_19950120T040000Z.png": img461,
	"gridrad_3d_v3_1_col_max_refl_19950120T050000Z.png": img462,
	"gridrad_3d_v3_1_col_max_refl_19950120T060000Z.png": img463,
	"gridrad_3d_v3_1_col_max_refl_19950120T070000Z.png": img464,
	"gridrad_3d_v3_1_col_max_refl_19950120T080000Z.png": img465,
	"gridrad_3d_v3_1_col_max_refl_19950120T090000Z.png": img466,
	"gridrad_3d_v3_1_col_max_refl_19950120T100000Z.png": img467,
	"gridrad_3d_v3_1_col_max_refl_19950120T110000Z.png": img468,
	"gridrad_3d_v3_1_col_max_refl_19950120T120000Z.png": img469,
	"gridrad_3d_v3_1_col_max_refl_19950120T130000Z.png": img470,
	"gridrad_3d_v3_1_col_max_refl_19950120T140000Z.png": img471,
	"gridrad_3d_v3_1_col_max_refl_19950120T150000Z.png": img472,
	"gridrad_3d_v3_1_col_max_refl_19950120T160000Z.png": img473,
	"gridrad_3d_v3_1_col_max_refl_19950120T170000Z.png": img474,
	"gridrad_3d_v3_1_col_max_refl_19950120T180000Z.png": img475,
	"gridrad_3d_v3_1_col_max_refl_19950120T190000Z.png": img476,
	"gridrad_3d_v3_1_col_max_refl_19950120T200000Z.png": img477,
	"gridrad_3d_v3_1_col_max_refl_19950120T210000Z.png": img478,
	"gridrad_3d_v3_1_col_max_refl_19950120T220000Z.png": img479,
	"gridrad_3d_v3_1_col_max_refl_19950120T230000Z.png": img480,
	"gridrad_3d_v3_1_col_max_refl_19950121T000000Z.png": img481,
	"gridrad_3d_v3_1_col_max_refl_19950121T010000Z.png": img482,
	"gridrad_3d_v3_1_col_max_refl_19950121T020000Z.png": img483,
	"gridrad_3d_v3_1_col_max_refl_19950121T030000Z.png": img484,
	"gridrad_3d_v3_1_col_max_refl_19950121T040000Z.png": img485,
	"gridrad_3d_v3_1_col_max_refl_19950121T050000Z.png": img486,
	"gridrad_3d_v3_1_col_max_refl_19950121T060000Z.png": img487,
	"gridrad_3d_v3_1_col_max_refl_19950121T070000Z.png": img488,
	"gridrad_3d_v3_1_col_max_refl_19950121T080000Z.png": img489,
	"gridrad_3d_v3_1_col_max_refl_19950121T090000Z.png": img490,
	"gridrad_3d_v3_1_col_max_refl_19950121T100000Z.png": img491,
	"gridrad_3d_v3_1_col_max_refl_19950121T110000Z.png": img492,
	"gridrad_3d_v3_1_col_max_refl_19950121T120000Z.png": img493,
	"gridrad_3d_v3_1_col_max_refl_19950121T130000Z.png": img494,
	"gridrad_3d_v3_1_col_max_refl_19950121T140000Z.png": img495,
	"gridrad_3d_v3_1_col_max_refl_19950121T150000Z.png": img496,
	"gridrad_3d_v3_1_col_max_refl_19950121T160000Z.png": img497,
	"gridrad_3d_v3_1_col_max_refl_19950121T170000Z.png": img498,
	"gridrad_3d_v3_1_col_max_refl_19950121T180000Z.png": img499,
	"gridrad_3d_v3_1_col_max_refl_19950121T190000Z.png": img500,
	"gridrad_3d_v3_1_col_max_refl_19950121T200000Z.png": img501,
	"gridrad_3d_v3_1_col_max_refl_19950121T210000Z.png": img502,
	"gridrad_3d_v3_1_col_max_refl_19950121T220000Z.png": img503,
	"gridrad_3d_v3_1_col_max_refl_19950121T230000Z.png": img504,
	"gridrad_3d_v3_1_col_max_refl_19950122T000000Z.png": img505,
	"gridrad_3d_v3_1_col_max_refl_19950122T010000Z.png": img506,
	"gridrad_3d_v3_1_col_max_refl_19950122T020000Z.png": img507,
	"gridrad_3d_v3_1_col_max_refl_19950122T030000Z.png": img508,
	"gridrad_3d_v3_1_col_max_refl_19950122T040000Z.png": img509,
	"gridrad_3d_v3_1_col_max_refl_19950122T050000Z.png": img510,
	"gridrad_3d_v3_1_col_max_refl_19950122T060000Z.png": img511,
	"gridrad_3d_v3_1_col_max_refl_19950122T070000Z.png": img512,
	"gridrad_3d_v3_1_col_max_refl_19950122T080000Z.png": img513,
	"gridrad_3d_v3_1_col_max_refl_19950122T090000Z.png": img514,
	"gridrad_3d_v3_1_col_max_refl_19950122T100000Z.png": img515,
	"gridrad_3d_v3_1_col_max_refl_19950122T110000Z.png": img516,
	"gridrad_3d_v3_1_col_max_refl_19950122T120000Z.png": img517,
	"gridrad_3d_v3_1_col_max_refl_19950122T130000Z.png": img518,
	"gridrad_3d_v3_1_col_max_refl_19950122T140000Z.png": img519,
	"gridrad_3d_v3_1_col_max_refl_19950122T150000Z.png": img520,
	"gridrad_3d_v3_1_col_max_refl_19950122T160000Z.png": img521,
	"gridrad_3d_v3_1_col_max_refl_19950122T170000Z.png": img522,
	"gridrad_3d_v3_1_col_max_refl_19950122T180000Z.png": img523,
	"gridrad_3d_v3_1_col_max_refl_19950122T190000Z.png": img524,
	"gridrad_3d_v3_1_col_max_refl_19950122T200000Z.png": img525,
	"gridrad_3d_v3_1_col_max_refl_19950122T210000Z.png": img526,
	"gridrad_3d_v3_1_col_max_refl_19950122T220000Z.png": img527,
	"gridrad_3d_v3_1_col_max_refl_19950122T230000Z.png": img528,
	"gridrad_3d_v3_1_col_max_refl_19950123T000000Z.png": img529,
	"gridrad_3d_v3_1_col_max_refl_19950123T010000Z.png": img530,
	"gridrad_3d_v3_1_col_max_refl_19950123T020000Z.png": img531,
	"gridrad_3d_v3_1_col_max_refl_19950123T030000Z.png": img532,
	"gridrad_3d_v3_1_col_max_refl_19950123T040000Z.png": img533,
	"gridrad_3d_v3_1_col_max_refl_19950123T050000Z.png": img534,
	"gridrad_3d_v3_1_col_max_refl_19950123T060000Z.png": img535,
	"gridrad_3d_v3_1_col_max_refl_19950123T070000Z.png": img536,
	"gridrad_3d_v3_1_col_max_refl_19950123T080000Z.png": img537,
	"gridrad_3d_v3_1_col_max_refl_19950123T090000Z.png": img538,
	"gridrad_3d_v3_1_col_max_refl_19950123T100000Z.png": img539,
	"gridrad_3d_v3_1_col_max_refl_19950123T110000Z.png": img540,
	"gridrad_3d_v3_1_col_max_refl_19950123T120000Z.png": img541,
	"gridrad_3d_v3_1_col_max_refl_19950123T130000Z.png": img542,
	"gridrad_3d_v3_1_col_max_refl_19950123T140000Z.png": img543,
	"gridrad_3d_v3_1_col_max_refl_19950123T150000Z.png": img544,
	"gridrad_3d_v3_1_col_max_refl_19950123T160000Z.png": img545,
	"gridrad_3d_v3_1_col_max_refl_19950123T170000Z.png": img546,
	"gridrad_3d_v3_1_col_max_refl_19950123T180000Z.png": img547,
	"gridrad_3d_v3_1_col_max_refl_19950123T190000Z.png": img548,
	"gridrad_3d_v3_1_col_max_refl_19950123T200000Z.png": img549,
	"gridrad_3d_v3_1_col_max_refl_19950123T210000Z.png": img550,
	"gridrad_3d_v3_1_col_max_refl_19950123T220000Z.png": img551,
	"gridrad_3d_v3_1_col_max_refl_19950123T230000Z.png": img552,
	"gridrad_3d_v3_1_col_max_refl_19950124T000000Z.png": img553,
	"gridrad_3d_v3_1_col_max_refl_19950124T010000Z.png": img554,
	"gridrad_3d_v3_1_col_max_refl_19950124T020000Z.png": img555,
	"gridrad_3d_v3_1_col_max_refl_19950124T030000Z.png": img556,
	"gridrad_3d_v3_1_col_max_refl_19950124T040000Z.png": img557,
	"gridrad_3d_v3_1_col_max_refl_19950124T050000Z.png": img558,
	"gridrad_3d_v3_1_col_max_refl_19950124T060000Z.png": img559,
	"gridrad_3d_v3_1_col_max_refl_19950124T070000Z.png": img560,
	"gridrad_3d_v3_1_col_max_refl_19950124T080000Z.png": img561,
	"gridrad_3d_v3_1_col_max_refl_19950124T090000Z.png": img562,
	"gridrad_3d_v3_1_col_max_refl_19950124T100000Z.png": img563,
	"gridrad_3d_v3_1_col_max_refl_19950124T110000Z.png": img564,
	"gridrad_3d_v3_1_col_max_refl_19950124T120000Z.png": img565,
	"gridrad_3d_v3_1_col_max_refl_19950124T130000Z.png": img566,
	"gridrad_3d_v3_1_col_max_refl_19950124T140000Z.png": img567,
	"gridrad_3d_v3_1_col_max_refl_19950124T150000Z.png": img568,
	"gridrad_3d_v3_1_col_max_refl_19950124T160000Z.png": img569,
	"gridrad_3d_v3_1_col_max_refl_19950124T170000Z.png": img570,
	"gridrad_3d_v3_1_col_max_refl_19950124T180000Z.png": img571,
	"gridrad_3d_v3_1_col_max_refl_19950124T190000Z.png": img572,
	"gridrad_3d_v3_1_col_max_refl_19950124T200000Z.png": img573,
	"gridrad_3d_v3_1_col_max_refl_19950124T210000Z.png": img574,
	"gridrad_3d_v3_1_col_max_refl_19950124T220000Z.png": img575,
	"gridrad_3d_v3_1_col_max_refl_19950124T230000Z.png": img576,
	"gridrad_3d_v3_1_col_max_refl_19950125T000000Z.png": img577,
	"gridrad_3d_v3_1_col_max_refl_19950125T010000Z.png": img578,
	"gridrad_3d_v3_1_col_max_refl_19950125T020000Z.png": img579,
	"gridrad_3d_v3_1_col_max_refl_19950125T030000Z.png": img580,
	"gridrad_3d_v3_1_col_max_refl_19950125T040000Z.png": img581,
	"gridrad_3d_v3_1_col_max_refl_19950125T050000Z.png": img582,
	"gridrad_3d_v3_1_col_max_refl_19950125T060000Z.png": img583,
	"gridrad_3d_v3_1_col_max_refl_19950125T070000Z.png": img584,
	"gridrad_3d_v3_1_col_max_refl_19950125T080000Z.png": img585,
	"gridrad_3d_v3_1_col_max_refl_19950125T090000Z.png": img586,
	"gridrad_3d_v3_1_col_max_refl_19950125T100000Z.png": img587,
	"gridrad_3d_v3_1_col_max_refl_19950125T110000Z.png": img588,
	"gridrad_3d_v3_1_col_max_refl_19950125T120000Z.png": img589,
	"gridrad_3d_v3_1_col_max_refl_19950125T130000Z.png": img590,
	"gridrad_3d_v3_1_col_max_refl_19950125T140000Z.png": img591,
	"gridrad_3d_v3_1_col_max_refl_19950125T150000Z.png": img592,
	"gridrad_3d_v3_1_col_max_refl_19950125T160000Z.png": img593,
	"gridrad_3d_v3_1_col_max_refl_19950125T170000Z.png": img594,
	"gridrad_3d_v3_1_col_max_refl_19950125T180000Z.png": img595,
	"gridrad_3d_v3_1_col_max_refl_19950125T190000Z.png": img596,
	"gridrad_3d_v3_1_col_max_refl_19950125T200000Z.png": img597,
	"gridrad_3d_v3_1_col_max_refl_19950125T210000Z.png": img598,
	"gridrad_3d_v3_1_col_max_refl_19950125T220000Z.png": img599,
	"gridrad_3d_v3_1_col_max_refl_19950125T230000Z.png": img600,
	"gridrad_3d_v3_1_col_max_refl_19950126T000000Z.png": img601,
	"gridrad_3d_v3_1_col_max_refl_19950126T010000Z.png": img602,
	"gridrad_3d_v3_1_col_max_refl_19950126T020000Z.png": img603,
	"gridrad_3d_v3_1_col_max_refl_19950126T030000Z.png": img604,
	"gridrad_3d_v3_1_col_max_refl_19950126T040000Z.png": img605,
	"gridrad_3d_v3_1_col_max_refl_19950126T050000Z.png": img606,
	"gridrad_3d_v3_1_col_max_refl_19950126T060000Z.png": img607,
	"gridrad_3d_v3_1_col_max_refl_19950126T070000Z.png": img608,
	"gridrad_3d_v3_1_col_max_refl_19950126T080000Z.png": img609,
	"gridrad_3d_v3_1_col_max_refl_19950126T090000Z.png": img610,
	"gridrad_3d_v3_1_col_max_refl_19950126T100000Z.png": img611,
	"gridrad_3d_v3_1_col_max_refl_19950126T110000Z.png": img612,
	"gridrad_3d_v3_1_col_max_refl_19950126T120000Z.png": img613,
	"gridrad_3d_v3_1_col_max_refl_19950126T130000Z.png": img614,
	"gridrad_3d_v3_1_col_max_refl_19950126T140000Z.png": img615,
	"gridrad_3d_v3_1_col_max_refl_19950126T150000Z.png": img616,
	"gridrad_3d_v3_1_col_max_refl_19950126T160000Z.png": img617,
	"gridrad_3d_v3_1_col_max_refl_19950126T170000Z.png": img618,
	"gridrad_3d_v3_1_col_max_refl_19950126T180000Z.png": img619,
	"gridrad_3d_v3_1_col_max_refl_19950126T190000Z.png": img620,
	"gridrad_3d_v3_1_col_max_refl_19950126T200000Z.png": img621,
	"gridrad_3d_v3_1_col_max_refl_19950126T210000Z.png": img622,
	"gridrad_3d_v3_1_col_max_refl_19950126T220000Z.png": img623,
	"gridrad_3d_v3_1_col_max_refl_19950126T230000Z.png": img624,
	"gridrad_3d_v3_1_col_max_refl_19950127T000000Z.png": img625,
	"gridrad_3d_v3_1_col_max_refl_19950127T010000Z.png": img626,
	"gridrad_3d_v3_1_col_max_refl_19950127T020000Z.png": img627,
	"gridrad_3d_v3_1_col_max_refl_19950127T030000Z.png": img628,
	"gridrad_3d_v3_1_col_max_refl_19950127T040000Z.png": img629,
	"gridrad_3d_v3_1_col_max_refl_19950127T050000Z.png": img630,
	"gridrad_3d_v3_1_col_max_refl_19950127T060000Z.png": img631,
	"gridrad_3d_v3_1_col_max_refl_19950127T070000Z.png": img632,
	"gridrad_3d_v3_1_col_max_refl_19950127T080000Z.png": img633,
	"gridrad_3d_v3_1_col_max_refl_19950127T090000Z.png": img634,
	"gridrad_3d_v3_1_col_max_refl_19950127T100000Z.png": img635,
	"gridrad_3d_v3_1_col_max_refl_19950127T110000Z.png": img636,
	"gridrad_3d_v3_1_col_max_refl_19950127T120000Z.png": img637,
	"gridrad_3d_v3_1_col_max_refl_19950127T130000Z.png": img638,
	"gridrad_3d_v3_1_col_max_refl_19950127T140000Z.png": img639,
	"gridrad_3d_v3_1_col_max_refl_19950127T150000Z.png": img640,
	"gridrad_3d_v3_1_col_max_refl_19950127T160000Z.png": img641,
	"gridrad_3d_v3_1_col_max_refl_19950127T170000Z.png": img642,
	"gridrad_3d_v3_1_col_max_refl_19950127T180000Z.png": img643,
	"gridrad_3d_v3_1_col_max_refl_19950127T190000Z.png": img644,
	"gridrad_3d_v3_1_col_max_refl_19950127T200000Z.png": img645,
	"gridrad_3d_v3_1_col_max_refl_19950127T210000Z.png": img646,
	"gridrad_3d_v3_1_col_max_refl_19950127T220000Z.png": img647,
	"gridrad_3d_v3_1_col_max_refl_19950127T230000Z.png": img648,
	"gridrad_3d_v3_1_col_max_refl_19950128T000000Z.png": img649,
	"gridrad_3d_v3_1_col_max_refl_19950128T010000Z.png": img650,
	"gridrad_3d_v3_1_col_max_refl_19950128T020000Z.png": img651,
	"gridrad_3d_v3_1_col_max_refl_19950128T030000Z.png": img652,
	"gridrad_3d_v3_1_col_max_refl_19950128T040000Z.png": img653,
	"gridrad_3d_v3_1_col_max_refl_19950128T050000Z.png": img654,
	"gridrad_3d_v3_1_col_max_refl_19950128T060000Z.png": img655,
	"gridrad_3d_v3_1_col_max_refl_19950128T070000Z.png": img656,
	"gridrad_3d_v3_1_col_max_refl_19950128T080000Z.png": img657,
	"gridrad_3d_v3_1_col_max_refl_19950128T090000Z.png": img658,
	"gridrad_3d_v3_1_col_max_refl_19950128T100000Z.png": img659,
	"gridrad_3d_v3_1_col_max_refl_19950128T110000Z.png": img660,
	"gridrad_3d_v3_1_col_max_refl_19950128T120000Z.png": img661,
	"gridrad_3d_v3_1_col_max_refl_19950128T130000Z.png": img662,
	"gridrad_3d_v3_1_col_max_refl_19950128T140000Z.png": img663,
	"gridrad_3d_v3_1_col_max_refl_19950128T150000Z.png": img664,
	"gridrad_3d_v3_1_col_max_refl_19950128T160000Z.png": img665,
	"gridrad_3d_v3_1_col_max_refl_19950128T170000Z.png": img666,
	"gridrad_3d_v3_1_col_max_refl_19950128T180000Z.png": img667,
	"gridrad_3d_v3_1_col_max_refl_19950128T190000Z.png": img668,
	"gridrad_3d_v3_1_col_max_refl_19950128T200000Z.png": img669,
	"gridrad_3d_v3_1_col_max_refl_19950128T210000Z.png": img670,
	"gridrad_3d_v3_1_col_max_refl_19950128T220000Z.png": img671,
	"gridrad_3d_v3_1_col_max_refl_19950128T230000Z.png": img672,
	"gridrad_3d_v3_1_col_max_refl_19950129T000000Z.png": img673,
	"gridrad_3d_v3_1_col_max_refl_19950129T010000Z.png": img674,
	"gridrad_3d_v3_1_col_max_refl_19950129T020000Z.png": img675,
	"gridrad_3d_v3_1_col_max_refl_19950129T030000Z.png": img676,
	"gridrad_3d_v3_1_col_max_refl_19950129T040000Z.png": img677,
	"gridrad_3d_v3_1_col_max_refl_19950129T050000Z.png": img678,
	"gridrad_3d_v3_1_col_max_refl_19950129T060000Z.png": img679,
	"gridrad_3d_v3_1_col_max_refl_19950129T070000Z.png": img680,
	"gridrad_3d_v3_1_col_max_refl_19950129T080000Z.png": img681,
	"gridrad_3d_v3_1_col_max_refl_19950129T090000Z.png": img682,
	"gridrad_3d_v3_1_col_max_refl_19950129T100000Z.png": img683,
	"gridrad_3d_v3_1_col_max_refl_19950129T110000Z.png": img684,
	"gridrad_3d_v3_1_col_max_refl_19950129T120000Z.png": img685,
	"gridrad_3d_v3_1_col_max_refl_19950129T130000Z.png": img686,
	"gridrad_3d_v3_1_col_max_refl_19950129T140000Z.png": img687,
	"gridrad_3d_v3_1_col_max_refl_19950129T150000Z.png": img688,
	"gridrad_3d_v3_1_col_max_refl_19950129T160000Z.png": img689,
	"gridrad_3d_v3_1_col_max_refl_19950129T170000Z.png": img690,
	"gridrad_3d_v3_1_col_max_refl_19950129T180000Z.png": img691,
	"gridrad_3d_v3_1_col_max_refl_19950129T190000Z.png": img692,
	"gridrad_3d_v3_1_col_max_refl_19950129T200000Z.png": img693,
	"gridrad_3d_v3_1_col_max_refl_19950129T210000Z.png": img694,
	"gridrad_3d_v3_1_col_max_refl_19950129T220000Z.png": img695,
	"gridrad_3d_v3_1_col_max_refl_19950129T230000Z.png": img696,
	"gridrad_3d_v3_1_col_max_refl_19950130T000000Z.png": img697,
	"gridrad_3d_v3_1_col_max_refl_19950130T010000Z.png": img698,
	"gridrad_3d_v3_1_col_max_refl_19950130T020000Z.png": img699,
	"gridrad_3d_v3_1_col_max_refl_19950130T030000Z.png": img700,
	"gridrad_3d_v3_1_col_max_refl_19950130T040000Z.png": img701,
	"gridrad_3d_v3_1_col_max_refl_19950130T050000Z.png": img702,
	"gridrad_3d_v3_1_col_max_refl_19950130T060000Z.png": img703,
	"gridrad_3d_v3_1_col_max_refl_19950130T070000Z.png": img704,
	"gridrad_3d_v3_1_col_max_refl_19950130T080000Z.png": img705,
	"gridrad_3d_v3_1_col_max_refl_19950130T090000Z.png": img706,
	"gridrad_3d_v3_1_col_max_refl_19950130T100000Z.png": img707,
	"gridrad_3d_v3_1_col_max_refl_19950130T110000Z.png": img708,
	"gridrad_3d_v3_1_col_max_refl_19950130T120000Z.png": img709,
	"gridrad_3d_v3_1_col_max_refl_19950130T130000Z.png": img710,
	"gridrad_3d_v3_1_col_max_refl_19950130T140000Z.png": img711,
	"gridrad_3d_v3_1_col_max_refl_19950130T150000Z.png": img712,
	"gridrad_3d_v3_1_col_max_refl_19950130T160000Z.png": img713,
	"gridrad_3d_v3_1_col_max_refl_19950130T170000Z.png": img714,
	"gridrad_3d_v3_1_col_max_refl_19950130T180000Z.png": img715,
	"gridrad_3d_v3_1_col_max_refl_19950130T190000Z.png": img716,
	"gridrad_3d_v3_1_col_max_refl_19950130T200000Z.png": img717,
	"gridrad_3d_v3_1_col_max_refl_19950130T210000Z.png": img718,
	"gridrad_3d_v3_1_col_max_refl_19950130T220000Z.png": img719,
	"gridrad_3d_v3_1_col_max_refl_19950130T230000Z.png": img720,
	"gridrad_3d_v3_1_col_max_refl_19950131T000000Z.png": img721,
	"gridrad_3d_v3_1_col_max_refl_19950131T010000Z.png": img722,
	"gridrad_3d_v3_1_col_max_refl_19950131T020000Z.png": img723,
	"gridrad_3d_v3_1_col_max_refl_19950131T030000Z.png": img724,
	"gridrad_3d_v3_1_col_max_refl_19950131T040000Z.png": img725,
	"gridrad_3d_v3_1_col_max_refl_19950131T050000Z.png": img726,
	"gridrad_3d_v3_1_col_max_refl_19950131T060000Z.png": img727,
	"gridrad_3d_v3_1_col_max_refl_19950131T070000Z.png": img728,
	"gridrad_3d_v3_1_col_max_refl_19950131T080000Z.png": img729,
	"gridrad_3d_v3_1_col_max_refl_19950131T090000Z.png": img730,
	"gridrad_3d_v3_1_col_max_refl_19950131T100000Z.png": img731,
	"gridrad_3d_v3_1_col_max_refl_19950131T110000Z.png": img732,
	"gridrad_3d_v3_1_col_max_refl_19950131T120000Z.png": img733,
	"gridrad_3d_v3_1_col_max_refl_19950131T130000Z.png": img734,
	"gridrad_3d_v3_1_col_max_refl_19950131T140000Z.png": img735,
	"gridrad_3d_v3_1_col_max_refl_19950131T150000Z.png": img736,
	"gridrad_3d_v3_1_col_max_refl_19950131T160000Z.png": img737,
	"gridrad_3d_v3_1_col_max_refl_19950131T170000Z.png": img738,
	"gridrad_3d_v3_1_col_max_refl_19950131T180000Z.png": img739,
	"gridrad_3d_v3_1_col_max_refl_19950131T190000Z.png": img740,
	"gridrad_3d_v3_1_col_max_refl_19950131T200000Z.png": img741,
	"gridrad_3d_v3_1_col_max_refl_19950131T210000Z.png": img742,
	"gridrad_3d_v3_1_col_max_refl_19950131T220000Z.png": img743,
	"gridrad_3d_v3_1_col_max_refl_19950131T230000Z.png": img744,
}


func serveImages(w http.ResponseWriter, r *http.Request) {
	img, ok := images[r.URL.Path]
	if !ok {
		http.Error(w, "image not found", http.StatusNotFound)
	}
	w.Write(img)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./images/"))))
	//http.HandleFunc("/", serveImages)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
