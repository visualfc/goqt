// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLocale_MeasurementSystem - QLocale::MeasurementSystem
type QLocale_MeasurementSystem uint32
const (
	QLocale_MetricSystem QLocale_MeasurementSystem = 0
	QLocale_ImperialSystem QLocale_MeasurementSystem = 1
)
//enum QLocale_NumberOption - QLocale::NumberOption
type QLocale_NumberOption uint32
const (
	QLocale_OmitGroupSeparator QLocale_NumberOption = 0x01
	QLocale_RejectGroupSeparator QLocale_NumberOption = 0x02
)
//enum QLocale_FormatType - QLocale::FormatType
type QLocale_FormatType uint32
const (
	QLocale_LongFormat QLocale_FormatType = 0
	QLocale_ShortFormat QLocale_FormatType = 1
	QLocale_NarrowFormat QLocale_FormatType = 2
)
//enum QLocale_Country - QLocale::Country
type QLocale_Country uint32
const (
	QLocale_AnyCountry QLocale_Country = 0
	QLocale_Afghanistan QLocale_Country = 1
	QLocale_Albania QLocale_Country = 2
	QLocale_Algeria QLocale_Country = 3
	QLocale_AmericanSamoa QLocale_Country = 4
	QLocale_Andorra QLocale_Country = 5
	QLocale_Angola QLocale_Country = 6
	QLocale_Anguilla QLocale_Country = 7
	QLocale_Antarctica QLocale_Country = 8
	QLocale_AntiguaAndBarbuda QLocale_Country = 9
	QLocale_Argentina QLocale_Country = 10
	QLocale_Armenia QLocale_Country = 11
	QLocale_Aruba QLocale_Country = 12
	QLocale_Australia QLocale_Country = 13
	QLocale_Austria QLocale_Country = 14
	QLocale_Azerbaijan QLocale_Country = 15
	QLocale_Bahamas QLocale_Country = 16
	QLocale_Bahrain QLocale_Country = 17
	QLocale_Bangladesh QLocale_Country = 18
	QLocale_Barbados QLocale_Country = 19
	QLocale_Belarus QLocale_Country = 20
	QLocale_Belgium QLocale_Country = 21
	QLocale_Belize QLocale_Country = 22
	QLocale_Benin QLocale_Country = 23
	QLocale_Bermuda QLocale_Country = 24
	QLocale_Bhutan QLocale_Country = 25
	QLocale_Bolivia QLocale_Country = 26
	QLocale_BosniaAndHerzegowina QLocale_Country = 27
	QLocale_Botswana QLocale_Country = 28
	QLocale_BouvetIsland QLocale_Country = 29
	QLocale_Brazil QLocale_Country = 30
	QLocale_BritishIndianOceanTerritory QLocale_Country = 31
	QLocale_BruneiDarussalam QLocale_Country = 32
	QLocale_Bulgaria QLocale_Country = 33
	QLocale_BurkinaFaso QLocale_Country = 34
	QLocale_Burundi QLocale_Country = 35
	QLocale_Cambodia QLocale_Country = 36
	QLocale_Cameroon QLocale_Country = 37
	QLocale_Canada QLocale_Country = 38
	QLocale_CapeVerde QLocale_Country = 39
	QLocale_CaymanIslands QLocale_Country = 40
	QLocale_CentralAfricanRepublic QLocale_Country = 41
	QLocale_Chad QLocale_Country = 42
	QLocale_Chile QLocale_Country = 43
	QLocale_China QLocale_Country = 44
	QLocale_ChristmasIsland QLocale_Country = 45
	QLocale_CocosIslands QLocale_Country = 46
	QLocale_Colombia QLocale_Country = 47
	QLocale_Comoros QLocale_Country = 48
	QLocale_DemocraticRepublicOfCongo QLocale_Country = 49
	QLocale_PeoplesRepublicOfCongo QLocale_Country = 50
	QLocale_CookIslands QLocale_Country = 51
	QLocale_CostaRica QLocale_Country = 52
	QLocale_IvoryCoast QLocale_Country = 53
	QLocale_Croatia QLocale_Country = 54
	QLocale_Cuba QLocale_Country = 55
	QLocale_Cyprus QLocale_Country = 56
	QLocale_CzechRepublic QLocale_Country = 57
	QLocale_Denmark QLocale_Country = 58
	QLocale_Djibouti QLocale_Country = 59
	QLocale_Dominica QLocale_Country = 60
	QLocale_DominicanRepublic QLocale_Country = 61
	QLocale_EastTimor QLocale_Country = 62
	QLocale_Ecuador QLocale_Country = 63
	QLocale_Egypt QLocale_Country = 64
	QLocale_ElSalvador QLocale_Country = 65
	QLocale_EquatorialGuinea QLocale_Country = 66
	QLocale_Eritrea QLocale_Country = 67
	QLocale_Estonia QLocale_Country = 68
	QLocale_Ethiopia QLocale_Country = 69
	QLocale_FalklandIslands QLocale_Country = 70
	QLocale_FaroeIslands QLocale_Country = 71
	QLocale_FijiCountry QLocale_Country = 72
	QLocale_Finland QLocale_Country = 73
	QLocale_France QLocale_Country = 74
	QLocale_MetropolitanFrance QLocale_Country = 75
	QLocale_FrenchGuiana QLocale_Country = 76
	QLocale_FrenchPolynesia QLocale_Country = 77
	QLocale_FrenchSouthernTerritories QLocale_Country = 78
	QLocale_Gabon QLocale_Country = 79
	QLocale_Gambia QLocale_Country = 80
	QLocale_Georgia QLocale_Country = 81
	QLocale_Germany QLocale_Country = 82
	QLocale_Ghana QLocale_Country = 83
	QLocale_Gibraltar QLocale_Country = 84
	QLocale_Greece QLocale_Country = 85
	QLocale_Greenland QLocale_Country = 86
	QLocale_Grenada QLocale_Country = 87
	QLocale_Guadeloupe QLocale_Country = 88
	QLocale_Guam QLocale_Country = 89
	QLocale_Guatemala QLocale_Country = 90
	QLocale_Guinea QLocale_Country = 91
	QLocale_GuineaBissau QLocale_Country = 92
	QLocale_Guyana QLocale_Country = 93
	QLocale_Haiti QLocale_Country = 94
	QLocale_HeardAndMcDonaldIslands QLocale_Country = 95
	QLocale_Honduras QLocale_Country = 96
	QLocale_HongKong QLocale_Country = 97
	QLocale_Hungary QLocale_Country = 98
	QLocale_Iceland QLocale_Country = 99
	QLocale_India QLocale_Country = 100
	QLocale_Indonesia QLocale_Country = 101
	QLocale_Iran QLocale_Country = 102
	QLocale_Iraq QLocale_Country = 103
	QLocale_Ireland QLocale_Country = 104
	QLocale_Israel QLocale_Country = 105
	QLocale_Italy QLocale_Country = 106
	QLocale_Jamaica QLocale_Country = 107
	QLocale_Japan QLocale_Country = 108
	QLocale_Jordan QLocale_Country = 109
	QLocale_Kazakhstan QLocale_Country = 110
	QLocale_Kenya QLocale_Country = 111
	QLocale_Kiribati QLocale_Country = 112
	QLocale_DemocraticRepublicOfKorea QLocale_Country = 113
	QLocale_RepublicOfKorea QLocale_Country = 114
	QLocale_Kuwait QLocale_Country = 115
	QLocale_Kyrgyzstan QLocale_Country = 116
	QLocale_Lao QLocale_Country = 117
	QLocale_Latvia QLocale_Country = 118
	QLocale_Lebanon QLocale_Country = 119
	QLocale_Lesotho QLocale_Country = 120
	QLocale_Liberia QLocale_Country = 121
	QLocale_LibyanArabJamahiriya QLocale_Country = 122
	QLocale_Liechtenstein QLocale_Country = 123
	QLocale_Lithuania QLocale_Country = 124
	QLocale_Luxembourg QLocale_Country = 125
	QLocale_Macau QLocale_Country = 126
	QLocale_Macedonia QLocale_Country = 127
	QLocale_Madagascar QLocale_Country = 128
	QLocale_Malawi QLocale_Country = 129
	QLocale_Malaysia QLocale_Country = 130
	QLocale_Maldives QLocale_Country = 131
	QLocale_Mali QLocale_Country = 132
	QLocale_Malta QLocale_Country = 133
	QLocale_MarshallIslands QLocale_Country = 134
	QLocale_Martinique QLocale_Country = 135
	QLocale_Mauritania QLocale_Country = 136
	QLocale_Mauritius QLocale_Country = 137
	QLocale_Mayotte QLocale_Country = 138
	QLocale_Mexico QLocale_Country = 139
	QLocale_Micronesia QLocale_Country = 140
	QLocale_Moldova QLocale_Country = 141
	QLocale_Monaco QLocale_Country = 142
	QLocale_Mongolia QLocale_Country = 143
	QLocale_Montserrat QLocale_Country = 144
	QLocale_Morocco QLocale_Country = 145
	QLocale_Mozambique QLocale_Country = 146
	QLocale_Myanmar QLocale_Country = 147
	QLocale_Namibia QLocale_Country = 148
	QLocale_NauruCountry QLocale_Country = 149
	QLocale_Nepal QLocale_Country = 150
	QLocale_Netherlands QLocale_Country = 151
	QLocale_NetherlandsAntilles QLocale_Country = 152
	QLocale_NewCaledonia QLocale_Country = 153
	QLocale_NewZealand QLocale_Country = 154
	QLocale_Nicaragua QLocale_Country = 155
	QLocale_Niger QLocale_Country = 156
	QLocale_Nigeria QLocale_Country = 157
	QLocale_Niue QLocale_Country = 158
	QLocale_NorfolkIsland QLocale_Country = 159
	QLocale_NorthernMarianaIslands QLocale_Country = 160
	QLocale_Norway QLocale_Country = 161
	QLocale_Oman QLocale_Country = 162
	QLocale_Pakistan QLocale_Country = 163
	QLocale_Palau QLocale_Country = 164
	QLocale_PalestinianTerritory QLocale_Country = 165
	QLocale_Panama QLocale_Country = 166
	QLocale_PapuaNewGuinea QLocale_Country = 167
	QLocale_Paraguay QLocale_Country = 168
	QLocale_Peru QLocale_Country = 169
	QLocale_Philippines QLocale_Country = 170
	QLocale_Pitcairn QLocale_Country = 171
	QLocale_Poland QLocale_Country = 172
	QLocale_Portugal QLocale_Country = 173
	QLocale_PuertoRico QLocale_Country = 174
	QLocale_Qatar QLocale_Country = 175
	QLocale_Reunion QLocale_Country = 176
	QLocale_Romania QLocale_Country = 177
	QLocale_RussianFederation QLocale_Country = 178
	QLocale_Rwanda QLocale_Country = 179
	QLocale_SaintKittsAndNevis QLocale_Country = 180
	QLocale_StLucia QLocale_Country = 181
	QLocale_StVincentAndTheGrenadines QLocale_Country = 182
	QLocale_Samoa QLocale_Country = 183
	QLocale_SanMarino QLocale_Country = 184
	QLocale_SaoTomeAndPrincipe QLocale_Country = 185
	QLocale_SaudiArabia QLocale_Country = 186
	QLocale_Senegal QLocale_Country = 187
	QLocale_Seychelles QLocale_Country = 188
	QLocale_SierraLeone QLocale_Country = 189
	QLocale_Singapore QLocale_Country = 190
	QLocale_Slovakia QLocale_Country = 191
	QLocale_Slovenia QLocale_Country = 192
	QLocale_SolomonIslands QLocale_Country = 193
	QLocale_Somalia QLocale_Country = 194
	QLocale_SouthAfrica QLocale_Country = 195
	QLocale_SouthGeorgiaAndTheSouthSandwichIslands QLocale_Country = 196
	QLocale_Spain QLocale_Country = 197
	QLocale_SriLanka QLocale_Country = 198
	QLocale_StHelena QLocale_Country = 199
	QLocale_StPierreAndMiquelon QLocale_Country = 200
	QLocale_Sudan QLocale_Country = 201
	QLocale_Suriname QLocale_Country = 202
	QLocale_SvalbardAndJanMayenIslands QLocale_Country = 203
	QLocale_Swaziland QLocale_Country = 204
	QLocale_Sweden QLocale_Country = 205
	QLocale_Switzerland QLocale_Country = 206
	QLocale_SyrianArabRepublic QLocale_Country = 207
	QLocale_Taiwan QLocale_Country = 208
	QLocale_Tajikistan QLocale_Country = 209
	QLocale_Tanzania QLocale_Country = 210
	QLocale_Thailand QLocale_Country = 211
	QLocale_Togo QLocale_Country = 212
	QLocale_Tokelau QLocale_Country = 213
	QLocale_TongaCountry QLocale_Country = 214
	QLocale_TrinidadAndTobago QLocale_Country = 215
	QLocale_Tunisia QLocale_Country = 216
	QLocale_Turkey QLocale_Country = 217
	QLocale_Turkmenistan QLocale_Country = 218
	QLocale_TurksAndCaicosIslands QLocale_Country = 219
	QLocale_Tuvalu QLocale_Country = 220
	QLocale_Uganda QLocale_Country = 221
	QLocale_Ukraine QLocale_Country = 222
	QLocale_UnitedArabEmirates QLocale_Country = 223
	QLocale_UnitedKingdom QLocale_Country = 224
	QLocale_UnitedStates QLocale_Country = 225
	QLocale_UnitedStatesMinorOutlyingIslands QLocale_Country = 226
	QLocale_Uruguay QLocale_Country = 227
	QLocale_Uzbekistan QLocale_Country = 228
	QLocale_Vanuatu QLocale_Country = 229
	QLocale_VaticanCityState QLocale_Country = 230
	QLocale_Venezuela QLocale_Country = 231
	QLocale_VietNam QLocale_Country = 232
	QLocale_BritishVirginIslands QLocale_Country = 233
	QLocale_USVirginIslands QLocale_Country = 234
	QLocale_WallisAndFutunaIslands QLocale_Country = 235
	QLocale_WesternSahara QLocale_Country = 236
	QLocale_Yemen QLocale_Country = 237
	QLocale_Yugoslavia QLocale_Country = 238
	QLocale_Zambia QLocale_Country = 239
	QLocale_Zimbabwe QLocale_Country = 240
	QLocale_SerbiaAndMontenegro QLocale_Country = 241
	QLocale_Montenegro QLocale_Country = 242
	QLocale_Serbia QLocale_Country = 243
	QLocale_SaintBarthelemy QLocale_Country = 244
	QLocale_SaintMartin QLocale_Country = 245
	QLocale_LatinAmericaAndTheCaribbean QLocale_Country = 246
	QLocale_LastCountry QLocale_Country = QLocale_LatinAmericaAndTheCaribbean
)
//enum QLocale_Language - QLocale::Language
type QLocale_Language uint32
const (
	QLocale_C QLocale_Language = 1
	QLocale_Abkhazian QLocale_Language = 2
	QLocale_Afan QLocale_Language = 3
	QLocale_Afar QLocale_Language = 4
	QLocale_Afrikaans QLocale_Language = 5
	QLocale_Albanian QLocale_Language = 6
	QLocale_Amharic QLocale_Language = 7
	QLocale_Arabic QLocale_Language = 8
	QLocale_Armenian QLocale_Language = 9
	QLocale_Assamese QLocale_Language = 10
	QLocale_Aymara QLocale_Language = 11
	QLocale_Azerbaijani QLocale_Language = 12
	QLocale_Bashkir QLocale_Language = 13
	QLocale_Basque QLocale_Language = 14
	QLocale_Bengali QLocale_Language = 15
	QLocale_Bhutani QLocale_Language = 16
	QLocale_Bihari QLocale_Language = 17
	QLocale_Bislama QLocale_Language = 18
	QLocale_Breton QLocale_Language = 19
	QLocale_Bulgarian QLocale_Language = 20
	QLocale_Burmese QLocale_Language = 21
	QLocale_Byelorussian QLocale_Language = 22
	QLocale_Cambodian QLocale_Language = 23
	QLocale_Catalan QLocale_Language = 24
	QLocale_Chinese QLocale_Language = 25
	QLocale_Corsican QLocale_Language = 26
	QLocale_Croatian QLocale_Language = 27
	QLocale_Czech QLocale_Language = 28
	QLocale_Danish QLocale_Language = 29
	QLocale_Dutch QLocale_Language = 30
	QLocale_English QLocale_Language = 31
	QLocale_Esperanto QLocale_Language = 32
	QLocale_Estonian QLocale_Language = 33
	QLocale_Faroese QLocale_Language = 34
	QLocale_FijiLanguage QLocale_Language = 35
	QLocale_Finnish QLocale_Language = 36
	QLocale_French QLocale_Language = 37
	QLocale_Frisian QLocale_Language = 38
	QLocale_Gaelic QLocale_Language = 39
	QLocale_Galician QLocale_Language = 40
	QLocale_Georgian QLocale_Language = 41
	QLocale_German QLocale_Language = 42
	QLocale_Greek QLocale_Language = 43
	QLocale_Greenlandic QLocale_Language = 44
	QLocale_Guarani QLocale_Language = 45
	QLocale_Gujarati QLocale_Language = 46
	QLocale_Hausa QLocale_Language = 47
	QLocale_Hebrew QLocale_Language = 48
	QLocale_Hindi QLocale_Language = 49
	QLocale_Hungarian QLocale_Language = 50
	QLocale_Icelandic QLocale_Language = 51
	QLocale_Indonesian QLocale_Language = 52
	QLocale_Interlingua QLocale_Language = 53
	QLocale_Interlingue QLocale_Language = 54
	QLocale_Inuktitut QLocale_Language = 55
	QLocale_Inupiak QLocale_Language = 56
	QLocale_Irish QLocale_Language = 57
	QLocale_Italian QLocale_Language = 58
	QLocale_Japanese QLocale_Language = 59
	QLocale_Javanese QLocale_Language = 60
	QLocale_Kannada QLocale_Language = 61
	QLocale_Kashmiri QLocale_Language = 62
	QLocale_Kazakh QLocale_Language = 63
	QLocale_Kinyarwanda QLocale_Language = 64
	QLocale_Kirghiz QLocale_Language = 65
	QLocale_Korean QLocale_Language = 66
	QLocale_Kurdish QLocale_Language = 67
	QLocale_Kurundi QLocale_Language = 68
	QLocale_Laothian QLocale_Language = 69
	QLocale_Latin QLocale_Language = 70
	QLocale_Latvian QLocale_Language = 71
	QLocale_Lingala QLocale_Language = 72
	QLocale_Lithuanian QLocale_Language = 73
	QLocale_Macedonian QLocale_Language = 74
	QLocale_Malagasy QLocale_Language = 75
	QLocale_Malay QLocale_Language = 76
	QLocale_Malayalam QLocale_Language = 77
	QLocale_Maltese QLocale_Language = 78
	QLocale_Maori QLocale_Language = 79
	QLocale_Marathi QLocale_Language = 80
	QLocale_Moldavian QLocale_Language = 81
	QLocale_Mongolian QLocale_Language = 82
	QLocale_NauruLanguage QLocale_Language = 83
	QLocale_Nepali QLocale_Language = 84
	QLocale_Norwegian QLocale_Language = 85
	QLocale_NorwegianBokmal QLocale_Language = QLocale_Norwegian
	QLocale_Occitan QLocale_Language = 86
	QLocale_Oriya QLocale_Language = 87
	QLocale_Pashto QLocale_Language = 88
	QLocale_Persian QLocale_Language = 89
	QLocale_Polish QLocale_Language = 90
	QLocale_Portuguese QLocale_Language = 91
	QLocale_Punjabi QLocale_Language = 92
	QLocale_Quechua QLocale_Language = 93
	QLocale_RhaetoRomance QLocale_Language = 94
	QLocale_Romanian QLocale_Language = 95
	QLocale_Russian QLocale_Language = 96
	QLocale_Samoan QLocale_Language = 97
	QLocale_Sangho QLocale_Language = 98
	QLocale_Sanskrit QLocale_Language = 99
	QLocale_Serbian QLocale_Language = 100
	QLocale_SerboCroatian QLocale_Language = 101
	QLocale_Sesotho QLocale_Language = 102
	QLocale_Setswana QLocale_Language = 103
	QLocale_Shona QLocale_Language = 104
	QLocale_Sindhi QLocale_Language = 105
	QLocale_Singhalese QLocale_Language = 106
	QLocale_Siswati QLocale_Language = 107
	QLocale_Slovak QLocale_Language = 108
	QLocale_Slovenian QLocale_Language = 109
	QLocale_Somali QLocale_Language = 110
	QLocale_Spanish QLocale_Language = 111
	QLocale_Sundanese QLocale_Language = 112
	QLocale_Swahili QLocale_Language = 113
	QLocale_Swedish QLocale_Language = 114
	QLocale_Tagalog QLocale_Language = 115
	QLocale_Tajik QLocale_Language = 116
	QLocale_Tamil QLocale_Language = 117
	QLocale_Tatar QLocale_Language = 118
	QLocale_Telugu QLocale_Language = 119
	QLocale_Thai QLocale_Language = 120
	QLocale_Tibetan QLocale_Language = 121
	QLocale_Tigrinya QLocale_Language = 122
	QLocale_TongaLanguage QLocale_Language = 123
	QLocale_Tsonga QLocale_Language = 124
	QLocale_Turkish QLocale_Language = 125
	QLocale_Turkmen QLocale_Language = 126
	QLocale_Twi QLocale_Language = 127
	QLocale_Uigur QLocale_Language = 128
	QLocale_Ukrainian QLocale_Language = 129
	QLocale_Urdu QLocale_Language = 130
	QLocale_Uzbek QLocale_Language = 131
	QLocale_Vietnamese QLocale_Language = 132
	QLocale_Volapuk QLocale_Language = 133
	QLocale_Welsh QLocale_Language = 134
	QLocale_Wolof QLocale_Language = 135
	QLocale_Xhosa QLocale_Language = 136
	QLocale_Yiddish QLocale_Language = 137
	QLocale_Yoruba QLocale_Language = 138
	QLocale_Zhuang QLocale_Language = 139
	QLocale_Zulu QLocale_Language = 140
	QLocale_NorwegianNynorsk QLocale_Language = 141
	QLocale_Nynorsk QLocale_Language = QLocale_NorwegianNynorsk
	QLocale_Bosnian QLocale_Language = 142
	QLocale_Divehi QLocale_Language = 143
	QLocale_Manx QLocale_Language = 144
	QLocale_Cornish QLocale_Language = 145
	QLocale_Akan QLocale_Language = 146
	QLocale_Konkani QLocale_Language = 147
	QLocale_Ga QLocale_Language = 148
	QLocale_Igbo QLocale_Language = 149
	QLocale_Kamba QLocale_Language = 150
	QLocale_Syriac QLocale_Language = 151
	QLocale_Blin QLocale_Language = 152
	QLocale_Geez QLocale_Language = 153
	QLocale_Koro QLocale_Language = 154
	QLocale_Sidamo QLocale_Language = 155
	QLocale_Atsam QLocale_Language = 156
	QLocale_Tigre QLocale_Language = 157
	QLocale_Jju QLocale_Language = 158
	QLocale_Friulian QLocale_Language = 159
	QLocale_Venda QLocale_Language = 160
	QLocale_Ewe QLocale_Language = 161
	QLocale_Walamo QLocale_Language = 162
	QLocale_Hawaiian QLocale_Language = 163
	QLocale_Tyap QLocale_Language = 164
	QLocale_Chewa QLocale_Language = 165
	QLocale_Filipino QLocale_Language = 166
	QLocale_SwissGerman QLocale_Language = 167
	QLocale_SichuanYi QLocale_Language = 168
	QLocale_Kpelle QLocale_Language = 169
	QLocale_LowGerman QLocale_Language = 170
	QLocale_SouthNdebele QLocale_Language = 171
	QLocale_NorthernSotho QLocale_Language = 172
	QLocale_NorthernSami QLocale_Language = 173
	QLocale_Taroko QLocale_Language = 174
	QLocale_Gusii QLocale_Language = 175
	QLocale_Taita QLocale_Language = 176
	QLocale_Fulah QLocale_Language = 177
	QLocale_Kikuyu QLocale_Language = 178
	QLocale_Samburu QLocale_Language = 179
	QLocale_Sena QLocale_Language = 180
	QLocale_NorthNdebele QLocale_Language = 181
	QLocale_Rombo QLocale_Language = 182
	QLocale_Tachelhit QLocale_Language = 183
	QLocale_Kabyle QLocale_Language = 184
	QLocale_Nyankole QLocale_Language = 185
	QLocale_Bena QLocale_Language = 186
	QLocale_Vunjo QLocale_Language = 187
	QLocale_Bambara QLocale_Language = 188
	QLocale_Embu QLocale_Language = 189
	QLocale_Cherokee QLocale_Language = 190
	QLocale_Morisyen QLocale_Language = 191
	QLocale_Makonde QLocale_Language = 192
	QLocale_Langi QLocale_Language = 193
	QLocale_Ganda QLocale_Language = 194
	QLocale_Bemba QLocale_Language = 195
	QLocale_Kabuverdianu QLocale_Language = 196
	QLocale_Meru QLocale_Language = 197
	QLocale_Kalenjin QLocale_Language = 198
	QLocale_Nama QLocale_Language = 199
	QLocale_Machame QLocale_Language = 200
	QLocale_Colognian QLocale_Language = 201
	QLocale_Masai QLocale_Language = 202
	QLocale_Soga QLocale_Language = 203
	QLocale_Luyia QLocale_Language = 204
	QLocale_Asu QLocale_Language = 205
	QLocale_Teso QLocale_Language = 206
	QLocale_Saho QLocale_Language = 207
	QLocale_KoyraChiini QLocale_Language = 208
	QLocale_Rwa QLocale_Language = 209
	QLocale_Luo QLocale_Language = 210
	QLocale_Chiga QLocale_Language = 211
	QLocale_CentralMoroccoTamazight QLocale_Language = 212
	QLocale_KoyraboroSenni QLocale_Language = 213
	QLocale_Shambala QLocale_Language = 214
	QLocale_LastLanguage QLocale_Language = QLocale_Shambala
)
//struct QLocale : QLocale
type QLocale struct {
	BaseDrv
}
//QLocale::QLocale()	
func NewLocale() *QLocale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),72000,72102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLocale{}
	_p.SetDriver(__rv,72000,true)
	return _p
} 
//QLocale::QLocale(QLocale const&)	
func NewLocaleCopy(other *QLocale) *QLocale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),72000,72103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLocale{}
	_p.SetDriver(__rv,72000,true)
	return _p
} 
//QLocale::QLocale(QString const&)	
func NewLocaleWithName(name string) *QLocale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),72000,72104,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLocale{}
	_p.SetDriver(__rv,72000,true)
	return _p
} 
//QLocale::QLocale(QLocale::Language,QLocale::Country)	
func NewLocaleWithLanguageCountry(language QLocale_Language,country QLocale_Country) *QLocale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),72000,72105,unsafe.Pointer(&language),unsafe.Pointer(&country),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLocale{}
	_p.SetDriver(__rv,72000,true)
	return _p
} 
//QLocale::amText()
func (q *QLocale) AmText() string {
	var __rv string
	q.Drv(72000,72106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::c()	
func QLocaleC() *QLocale {
	var __rv uintptr
	DirectQtDrv(nil,72000,72107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QLocale::c()
func (q *QLocale) C() *QLocale {
	var __rv uintptr
	q.Drv(72000,72107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QLocale::countriesForLanguage(QLocale::Language)	
func QLocaleCountriesForLanguage(lang QLocale_Language) []QLocale_Country {
	var __rv []QLocale_Country
	DirectQtDrv(nil,72000,72108,unsafe.Pointer(&lang),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::countriesForLanguage(QLocale::Language)
func (q *QLocale) CountriesForLanguage(lang QLocale_Language) []QLocale_Country {
	var __rv []QLocale_Country
	q.Drv(72000,72108,unsafe.Pointer(&lang),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::country()
func (q *QLocale) Country() QLocale_Country {
	var __rv QLocale_Country
	q.Drv(72000,72109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::countryToString(QLocale::Country)	
func QLocaleCountryToString(country QLocale_Country) string {
	var __rv string
	DirectQtDrv(nil,72000,72110,unsafe.Pointer(&country),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::countryToString(QLocale::Country)
func (q *QLocale) CountryToString(country QLocale_Country) string {
	var __rv string
	q.Drv(72000,72110,unsafe.Pointer(&country),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dateFormat()
func (q *QLocale) DateFormat() string {
	var __rv string
	q.Drv(72000,72111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dateFormat(QLocale::FormatType)
func (q *QLocale) DateFormatWithFormatType(format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72112,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dateTimeFormat()
func (q *QLocale) DateTimeFormat() string {
	var __rv string
	q.Drv(72000,72113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dateTimeFormat(QLocale::FormatType)
func (q *QLocale) DateTimeFormatWithFormatType(format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72114,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dayName(int)
func (q *QLocale) DayName(value int) string {
	var __rv string
	q.Drv(72000,72115,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::dayName(int,QLocale::FormatType)
func (q *QLocale) DayNameWithIntFormatType(value2 int,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72116,unsafe.Pointer(&value2),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::decimalPoint()
func (q *QLocale) DecimalPoint() rune {
	var __rv rune
	q.Drv(72000,72117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::exponential()
func (q *QLocale) Exponential() rune {
	var __rv rune
	q.Drv(72000,72118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::groupSeparator()
func (q *QLocale) GroupSeparator() rune {
	var __rv rune
	q.Drv(72000,72119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::language()
func (q *QLocale) Language() QLocale_Language {
	var __rv QLocale_Language
	q.Drv(72000,72120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::languageToString(QLocale::Language)	
func QLocaleLanguageToString(language QLocale_Language) string {
	var __rv string
	DirectQtDrv(nil,72000,72121,unsafe.Pointer(&language),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::languageToString(QLocale::Language)
func (q *QLocale) LanguageToString(language QLocale_Language) string {
	var __rv string
	q.Drv(72000,72121,unsafe.Pointer(&language),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::measurementSystem()
func (q *QLocale) MeasurementSystem() QLocale_MeasurementSystem {
	var __rv QLocale_MeasurementSystem
	q.Drv(72000,72122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::monthName(int)
func (q *QLocale) MonthName(value int) string {
	var __rv string
	q.Drv(72000,72123,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::monthName(int,QLocale::FormatType)
func (q *QLocale) MonthNameWithIntFormatType(value2 int,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72124,unsafe.Pointer(&value2),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::name()
func (q *QLocale) Name() string {
	var __rv string
	q.Drv(72000,72125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::negativeSign()
func (q *QLocale) NegativeSign() rune {
	var __rv rune
	q.Drv(72000,72126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::numberOptions()
func (q *QLocale) NumberOptions() QLocale_NumberOption {
	var __rv QLocale_NumberOption
	q.Drv(72000,72127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::percent()
func (q *QLocale) Percent() rune {
	var __rv rune
	q.Drv(72000,72128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::pmText()
func (q *QLocale) PmText() string {
	var __rv string
	q.Drv(72000,72129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::positiveSign()
func (q *QLocale) PositiveSign() rune {
	var __rv rune
	q.Drv(72000,72130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::setDefault(QLocale const&)	
func QLocaleSetDefault(locale *QLocale)  {
	DirectQtDrv(nil,72000,72131,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLocale::setDefault(QLocale const&)
func (q *QLocale) SetDefault(locale *QLocale)  {
	q.Drv(72000,72131,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLocale::setNumberOptions(QFlags<QLocale::NumberOption>)
func (q *QLocale) SetNumberOptions(options QLocale_NumberOption)  {
	q.Drv(72000,72132,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLocale::standaloneDayName(int)
func (q *QLocale) StandaloneDayName(value int) string {
	var __rv string
	q.Drv(72000,72133,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::standaloneDayName(int,QLocale::FormatType)
func (q *QLocale) StandaloneDayNameWithIntFormatType(value2 int,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72134,unsafe.Pointer(&value2),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::standaloneMonthName(int)
func (q *QLocale) StandaloneMonthName(value int) string {
	var __rv string
	q.Drv(72000,72135,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::standaloneMonthName(int,QLocale::FormatType)
func (q *QLocale) StandaloneMonthNameWithIntFormatType(value2 int,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72136,unsafe.Pointer(&value2),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::system()	
func QLocaleSystem() *QLocale {
	var __rv uintptr
	DirectQtDrv(nil,72000,72137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QLocale::system()
func (q *QLocale) System() *QLocale {
	var __rv uintptr
	q.Drv(72000,72137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QLocale::textDirection()
func (q *QLocale) TextDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(72000,72138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::timeFormat()
func (q *QLocale) TimeFormat() string {
	var __rv string
	q.Drv(72000,72139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::timeFormat(QLocale::FormatType)
func (q *QLocale) TimeFormatWithFormatType(format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72140,unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toDate(QString const&)
func (q *QLocale) ToDate(string string) *QDate {
	var __rv uintptr
	q.Drv(72000,72141,unsafe.Pointer(&string),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QLocale::toDate(QString const&,QLocale::FormatType)
func (q *QLocale) ToDateWithStringFormatType(string string,value2 QLocale_FormatType) *QDate {
	var __rv uintptr
	q.Drv(72000,72142,unsafe.Pointer(&string),unsafe.Pointer(&value2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QLocale::toDate(QString const&,QString const&)
func (q *QLocale) ToDateWithStringFormat(string string,format string) *QDate {
	var __rv uintptr
	q.Drv(72000,72143,unsafe.Pointer(&string),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QLocale::toDateTime(QString const&)
func (q *QLocale) ToDateTime(string string) *QDateTime {
	var __rv uintptr
	q.Drv(72000,72144,unsafe.Pointer(&string),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QLocale::toDateTime(QString const&,QLocale::FormatType)
func (q *QLocale) ToDateTimeWithStringFormatType(string string,format QLocale_FormatType) *QDateTime {
	var __rv uintptr
	q.Drv(72000,72145,unsafe.Pointer(&string),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QLocale::toDateTime(QString const&,QString const&)
func (q *QLocale) ToDateTimeWithStringFormat(string string,format string) *QDateTime {
	var __rv uintptr
	q.Drv(72000,72146,unsafe.Pointer(&string),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QLocale::toDouble(QString const&)
func (q *QLocale) ToDouble(s string) float64 {
	var __rv float64
	q.Drv(72000,72147,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toDouble(QString const&,bool*)
func (q *QLocale) ToDoubleWithTextOk(s string,ok *bool) float64 {
	var __rv float64
	q.Drv(72000,72148,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toFloat(QString const&)
func (q *QLocale) ToFloat(s string) float32 {
	var __rv float32
	q.Drv(72000,72149,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toFloat(QString const&,bool*)
func (q *QLocale) ToFloatWithTextOk(s string,ok *bool) float32 {
	var __rv float32
	q.Drv(72000,72150,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toInt(QString const&)
func (q *QLocale) ToInt(s string) int {
	var __rv int
	q.Drv(72000,72151,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toInt(QString const&,bool*,int)
func (q *QLocale) ToIntWithTextOkBase(s string,ok *bool) int {
	var __rv int
	q.Drv(72000,72152,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toLongLong(QString const&)
func (q *QLocale) ToLongLong(s string) int64 {
	var __rv int64
	q.Drv(72000,72153,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toLongLong(QString const&,bool*,int)
func (q *QLocale) ToLongLongWithTextOkBase(s string,ok *bool) int64 {
	var __rv int64
	q.Drv(72000,72154,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toShort(QString const&)
func (q *QLocale) ToShort(s string) int16 {
	var __rv int16
	q.Drv(72000,72155,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toShort(QString const&,bool*,int)
func (q *QLocale) ToShortWithTextOkBase(s string,ok *bool) int16 {
	var __rv int16
	q.Drv(72000,72156,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDate const&)
func (q *QLocale) ToString(date *QDate) string {
	var __rv string
	q.Drv(72000,72157,Native(date),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDateTime const&)
func (q *QLocale) ToStringWithDatetime(dateTime *QDateTime) string {
	var __rv string
	q.Drv(72000,72158,Native(dateTime),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QTime const&)
func (q *QLocale) ToStringWithTime(time *QTime) string {
	var __rv string
	q.Drv(72000,72159,Native(time),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(double)
func (q *QLocale) ToStringFWithFloat64(i float64) string {
	var __rv string
	q.Drv(72000,72160,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(float)
func (q *QLocale) ToStringFWithFloat32(i float32) string {
	var __rv string
	q.Drv(72000,72161,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(int)
func (q *QLocale) ToStringWithInt(i int) string {
	var __rv string
	q.Drv(72000,72162,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(qint64)
func (q *QLocale) ToStringWithInt64(i int64) string {
	var __rv string
	q.Drv(72000,72163,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(short)
func (q *QLocale) ToStringWithInt16(i int16) string {
	var __rv string
	q.Drv(72000,72164,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(unsigned int)
func (q *QLocale) ToStringWithUint(i uint) string {
	var __rv string
	q.Drv(72000,72165,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(unsigned short)
func (q *QLocale) ToStringWithUint16(i uint16) string {
	var __rv string
	q.Drv(72000,72166,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDate const&,QLocale::FormatType)
func (q *QLocale) ToStringWithDateFormatType(date *QDate,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72167,Native(date),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDate const&,QString const&)
func (q *QLocale) ToStringWithDateFormatstr(date *QDate,formatStr string) string {
	var __rv string
	q.Drv(72000,72168,Native(date),unsafe.Pointer(&formatStr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDateTime const&,QLocale::FormatType)
func (q *QLocale) ToStringWithDatetimeFormatType(dateTime *QDateTime,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72169,Native(dateTime),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QDateTime const&,QString const&)
func (q *QLocale) ToStringWithDatetimeFormat(dateTime *QDateTime,format string) string {
	var __rv string
	q.Drv(72000,72170,Native(dateTime),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QTime const&,QLocale::FormatType)
func (q *QLocale) ToStringWithTimeFormatType(time *QTime,format QLocale_FormatType) string {
	var __rv string
	q.Drv(72000,72171,Native(time),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(QTime const&,QString const&)
func (q *QLocale) ToStringWithTimeFormatstr(time *QTime,formatStr string) string {
	var __rv string
	q.Drv(72000,72172,Native(time),unsafe.Pointer(&formatStr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(double,char,int)
func (q *QLocale) ToStringFWithFloat64FmtPrec(i float64,f byte,prec int) string {
	var __rv string
	q.Drv(72000,72173,unsafe.Pointer(&i),unsafe.Pointer(&f),unsafe.Pointer(&prec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toString(float,char,int)
func (q *QLocale) ToStringFWithFloat32FmtPrec(i float32,f byte,prec int) string {
	var __rv string
	q.Drv(72000,72174,unsafe.Pointer(&i),unsafe.Pointer(&f),unsafe.Pointer(&prec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toTime(QString const&)
func (q *QLocale) ToTime(string string) *QTime {
	var __rv uintptr
	q.Drv(72000,72175,unsafe.Pointer(&string),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QLocale::toTime(QString const&,QLocale::FormatType)
func (q *QLocale) ToTimeWithStringFormatType(string string,value2 QLocale_FormatType) *QTime {
	var __rv uintptr
	q.Drv(72000,72176,unsafe.Pointer(&string),unsafe.Pointer(&value2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QLocale::toTime(QString const&,QString const&)
func (q *QLocale) ToTimeWithStringFormat(string string,format string) *QTime {
	var __rv uintptr
	q.Drv(72000,72177,unsafe.Pointer(&string),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTime{}
	_rp.SetDriver(__rv,170000,true)
	return _rp
}	
//QLocale::toUInt(QString const&)
func (q *QLocale) ToUInt(s string) uint {
	var __rv uint
	q.Drv(72000,72178,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toUInt(QString const&,bool*,int)
func (q *QLocale) ToUIntWithTextOkBase(s string,ok *bool) uint {
	var __rv uint
	q.Drv(72000,72179,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toULongLong(QString const&)
func (q *QLocale) ToULongLong(s string) int64 {
	var __rv int64
	q.Drv(72000,72180,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toULongLong(QString const&,bool*,int)
func (q *QLocale) ToULongLongWithTextOkBase(s string,ok *bool) int64 {
	var __rv int64
	q.Drv(72000,72181,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toUShort(QString const&)
func (q *QLocale) ToUShort(s string) uint16 {
	var __rv uint16
	q.Drv(72000,72182,unsafe.Pointer(&s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::toUShort(QString const&,bool*,int)
func (q *QLocale) ToUShortWithTextOkBase(s string,ok *bool) uint16 {
	var __rv uint16
	q.Drv(72000,72183,unsafe.Pointer(&s),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLocale::zeroDigit()
func (q *QLocale) ZeroDigit() rune {
	var __rv rune
	q.Drv(72000,72184,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
