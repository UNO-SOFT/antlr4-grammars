// Generated from UnicodeClasses.g4 by ANTLR 4.7.

package unicodeclasses

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 31, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 3, 2, 9, 584, 2, 99, 124, 183, 183, 225, 248, 250, 257, 259,
	259, 261, 261, 263, 263, 265, 265, 267, 267, 269, 269, 271, 271, 273, 273,
	275, 275, 277, 277, 279, 279, 281, 281, 283, 283, 285, 285, 287, 287, 289,
	289, 291, 291, 293, 293, 295, 295, 297, 297, 299, 299, 301, 301, 303, 303,
	305, 305, 307, 307, 309, 309, 311, 311, 313, 314, 316, 316, 318, 318, 320,
	320, 322, 322, 324, 324, 326, 326, 328, 328, 330, 331, 333, 333, 335, 335,
	337, 337, 339, 339, 341, 341, 343, 343, 345, 345, 347, 347, 349, 349, 351,
	351, 353, 353, 355, 355, 357, 357, 359, 359, 361, 361, 363, 363, 365, 365,
	367, 367, 369, 369, 371, 371, 373, 373, 375, 375, 377, 377, 380, 380, 382,
	382, 384, 386, 389, 389, 391, 391, 394, 394, 398, 399, 404, 404, 407, 407,
	411, 413, 416, 416, 419, 419, 421, 421, 423, 423, 426, 426, 428, 429, 431,
	431, 434, 434, 438, 438, 440, 440, 443, 444, 447, 449, 456, 456, 459, 459,
	462, 462, 464, 464, 466, 466, 468, 468, 470, 470, 472, 472, 474, 474, 476,
	476, 478, 479, 481, 481, 483, 483, 485, 485, 487, 487, 489, 489, 491, 491,
	493, 493, 495, 495, 497, 498, 501, 501, 503, 503, 507, 507, 509, 509, 511,
	511, 513, 513, 515, 515, 517, 517, 519, 519, 521, 521, 523, 523, 525, 525,
	527, 527, 529, 529, 531, 531, 533, 533, 535, 535, 537, 537, 539, 539, 541,
	541, 543, 543, 545, 545, 547, 547, 549, 549, 551, 551, 553, 553, 555, 555,
	557, 557, 559, 559, 561, 561, 563, 563, 565, 571, 574, 574, 577, 578, 580,
	580, 585, 585, 587, 587, 589, 589, 591, 591, 593, 661, 663, 689, 883, 883,
	885, 885, 889, 889, 893, 895, 914, 914, 942, 976, 978, 979, 983, 985, 987,
	987, 989, 989, 991, 991, 993, 993, 995, 995, 997, 997, 999, 999, 1001,
	1001, 1003, 1003, 1005, 1005, 1007, 1007, 1009, 1013, 1015, 1015, 1018,
	1018, 1021, 1022, 1074, 1121, 1123, 1123, 1125, 1125, 1127, 1127, 1129,
	1129, 1131, 1131, 1133, 1133, 1135, 1135, 1137, 1137, 1139, 1139, 1141,
	1141, 1143, 1143, 1145, 1145, 1147, 1147, 1149, 1149, 1151, 1151, 1153,
	1153, 1155, 1155, 1165, 1165, 1167, 1167, 1169, 1169, 1171, 1171, 1173,
	1173, 1175, 1175, 1177, 1177, 1179, 1179, 1181, 1181, 1183, 1183, 1185,
	1185, 1187, 1187, 1189, 1189, 1191, 1191, 1193, 1193, 1195, 1195, 1197,
	1197, 1199, 1199, 1201, 1201, 1203, 1203, 1205, 1205, 1207, 1207, 1209,
	1209, 1211, 1211, 1213, 1213, 1215, 1215, 1217, 1217, 1220, 1220, 1222,
	1222, 1224, 1224, 1226, 1226, 1228, 1228, 1230, 1230, 1232, 1233, 1235,
	1235, 1237, 1237, 1239, 1239, 1241, 1241, 1243, 1243, 1245, 1245, 1247,
	1247, 1249, 1249, 1251, 1251, 1253, 1253, 1255, 1255, 1257, 1257, 1259,
	1259, 1261, 1261, 1263, 1263, 1265, 1265, 1267, 1267, 1269, 1269, 1271,
	1271, 1273, 1273, 1275, 1275, 1277, 1277, 1279, 1279, 1281, 1281, 1283,
	1283, 1285, 1285, 1287, 1287, 1289, 1289, 1291, 1291, 1293, 1293, 1295,
	1295, 1297, 1297, 1299, 1299, 1301, 1301, 1303, 1303, 1305, 1305, 1307,
	1307, 1309, 1309, 1311, 1311, 1313, 1313, 1315, 1315, 1317, 1317, 1319,
	1319, 1321, 1321, 1379, 1417, 7426, 7469, 7533, 7545, 7547, 7580, 7683,
	7683, 7685, 7685, 7687, 7687, 7689, 7689, 7691, 7691, 7693, 7693, 7695,
	7695, 7697, 7697, 7699, 7699, 7701, 7701, 7703, 7703, 7705, 7705, 7707,
	7707, 7709, 7709, 7711, 7711, 7713, 7713, 7715, 7715, 7717, 7717, 7719,
	7719, 7721, 7721, 7723, 7723, 7725, 7725, 7727, 7727, 7729, 7729, 7731,
	7731, 7733, 7733, 7735, 7735, 7737, 7737, 7739, 7739, 7741, 7741, 7743,
	7743, 7745, 7745, 7747, 7747, 7749, 7749, 7751, 7751, 7753, 7753, 7755,
	7755, 7757, 7757, 7759, 7759, 7761, 7761, 7763, 7763, 7765, 7765, 7767,
	7767, 7769, 7769, 7771, 7771, 7773, 7773, 7775, 7775, 7777, 7777, 7779,
	7779, 7781, 7781, 7783, 7783, 7785, 7785, 7787, 7787, 7789, 7789, 7791,
	7791, 7793, 7793, 7795, 7795, 7797, 7797, 7799, 7799, 7801, 7801, 7803,
	7803, 7805, 7805, 7807, 7807, 7809, 7809, 7811, 7811, 7813, 7813, 7815,
	7815, 7817, 7817, 7819, 7819, 7821, 7821, 7823, 7823, 7825, 7825, 7827,
	7827, 7829, 7829, 7831, 7839, 7841, 7841, 7843, 7843, 7845, 7845, 7847,
	7847, 7849, 7849, 7851, 7851, 7853, 7853, 7855, 7855, 7857, 7857, 7859,
	7859, 7861, 7861, 7863, 7863, 7865, 7865, 7867, 7867, 7869, 7869, 7871,
	7871, 7873, 7873, 7875, 7875, 7877, 7877, 7879, 7879, 7881, 7881, 7883,
	7883, 7885, 7885, 7887, 7887, 7889, 7889, 7891, 7891, 7893, 7893, 7895,
	7895, 7897, 7897, 7899, 7899, 7901, 7901, 7903, 7903, 7905, 7905, 7907,
	7907, 7909, 7909, 7911, 7911, 7913, 7913, 7915, 7915, 7917, 7917, 7919,
	7919, 7921, 7921, 7923, 7923, 7925, 7925, 7927, 7927, 7929, 7929, 7931,
	7931, 7933, 7933, 7935, 7935, 7937, 7945, 7954, 7959, 7970, 7977, 7986,
	7993, 8002, 8007, 8018, 8025, 8034, 8041, 8050, 8063, 8066, 8073, 8082,
	8089, 8098, 8105, 8114, 8118, 8120, 8121, 8128, 8128, 8132, 8134, 8136,
	8137, 8146, 8149, 8152, 8153, 8162, 8169, 8180, 8182, 8184, 8185, 8460,
	8460, 8464, 8465, 8469, 8469, 8497, 8497, 8502, 8502, 8507, 8507, 8510,
	8511, 8520, 8523, 8528, 8528, 8582, 8582, 11314, 11360, 11363, 11363, 11367,
	11368, 11370, 11370, 11372, 11372, 11374, 11374, 11379, 11379, 11381, 11382,
	11384, 11389, 11395, 11395, 11397, 11397, 11399, 11399, 11401, 11401, 11403,
	11403, 11405, 11405, 11407, 11407, 11409, 11409, 11411, 11411, 11413, 11413,
	11415, 11415, 11417, 11417, 11419, 11419, 11421, 11421, 11423, 11423, 11425,
	11425, 11427, 11427, 11429, 11429, 11431, 11431, 11433, 11433, 11435, 11435,
	11437, 11437, 11439, 11439, 11441, 11441, 11443, 11443, 11445, 11445, 11447,
	11447, 11449, 11449, 11451, 11451, 11453, 11453, 11455, 11455, 11457, 11457,
	11459, 11459, 11461, 11461, 11463, 11463, 11465, 11465, 11467, 11467, 11469,
	11469, 11471, 11471, 11473, 11473, 11475, 11475, 11477, 11477, 11479, 11479,
	11481, 11481, 11483, 11483, 11485, 11485, 11487, 11487, 11489, 11489, 11491,
	11491, 11493, 11494, 11502, 11502, 11504, 11504, 11509, 11509, 11522, 11559,
	11561, 11561, 11567, 11567, 42563, 42563, 42565, 42565, 42567, 42567, 42569,
	42569, 42571, 42571, 42573, 42573, 42575, 42575, 42577, 42577, 42579, 42579,
	42581, 42581, 42583, 42583, 42585, 42585, 42587, 42587, 42589, 42589, 42591,
	42591, 42593, 42593, 42595, 42595, 42597, 42597, 42599, 42599, 42601, 42601,
	42603, 42603, 42605, 42605, 42607, 42607, 42627, 42627, 42629, 42629, 42631,
	42631, 42633, 42633, 42635, 42635, 42637, 42637, 42639, 42639, 42641, 42641,
	42643, 42643, 42645, 42645, 42647, 42647, 42649, 42649, 42789, 42789, 42791,
	42791, 42793, 42793, 42795, 42795, 42797, 42797, 42799, 42799, 42801, 42803,
	42805, 42805, 42807, 42807, 42809, 42809, 42811, 42811, 42813, 42813, 42815,
	42815, 42817, 42817, 42819, 42819, 42821, 42821, 42823, 42823, 42825, 42825,
	42827, 42827, 42829, 42829, 42831, 42831, 42833, 42833, 42835, 42835, 42837,
	42837, 42839, 42839, 42841, 42841, 42843, 42843, 42845, 42845, 42847, 42847,
	42849, 42849, 42851, 42851, 42853, 42853, 42855, 42855, 42857, 42857, 42859,
	42859, 42861, 42861, 42863, 42863, 42865, 42865, 42867, 42874, 42876, 42876,
	42878, 42878, 42881, 42881, 42883, 42883, 42885, 42885, 42887, 42887, 42889,
	42889, 42894, 42894, 42896, 42896, 42899, 42899, 42901, 42901, 42915, 42915,
	42917, 42917, 42919, 42919, 42921, 42921, 42923, 42923, 43004, 43004, 64258,
	64264, 64277, 64281, 65347, 65372, 53, 2, 690, 707, 712, 723, 738, 742,
	750, 750, 752, 752, 886, 886, 892, 892, 1371, 1371, 1602, 1602, 1767, 1768,
	2038, 2039, 2044, 2044, 2076, 2076, 2086, 2086, 2090, 2090, 2419, 2419,
	3656, 3656, 3784, 3784, 4350, 4350, 6105, 6105, 6213, 6213, 6825, 6825,
	7290, 7295, 7470, 7532, 7546, 7546, 7581, 7617, 8307, 8307, 8321, 8321,
	8338, 8350, 11390, 11391, 11633, 11633, 11825, 11825, 12295, 12295, 12339,
	12343, 12349, 12349, 12447, 12448, 12542, 12544, 40983, 40983, 42234, 42239,
	42510, 42510, 42625, 42625, 42777, 42785, 42866, 42866, 42890, 42890, 43002,
	43003, 43473, 43473, 43634, 43634, 43743, 43743, 43765, 43766, 65394, 65394,
	65440, 65441, 291, 2, 172, 172, 188, 188, 445, 445, 450, 453, 662, 662,
	1490, 1516, 1522, 1524, 1570, 1601, 1603, 1612, 1648, 1649, 1651, 1749,
	1751, 1751, 1776, 1777, 1788, 1790, 1793, 1793, 1810, 1810, 1812, 1841,
	1871, 1959, 1971, 1971, 1996, 2028, 2050, 2071, 2114, 2138, 2210, 2210,
	2212, 2222, 2310, 2363, 2367, 2367, 2386, 2386, 2394, 2403, 2420, 2425,
	2427, 2433, 2439, 2446, 2449, 2450, 2453, 2474, 2476, 2482, 2484, 2484,
	2488, 2491, 2495, 2495, 2512, 2512, 2526, 2527, 2529, 2531, 2546, 2547,
	2567, 2572, 2577, 2578, 2581, 2602, 2604, 2610, 2612, 2613, 2615, 2616,
	2618, 2619, 2651, 2654, 2656, 2656, 2676, 2678, 2695, 2703, 2705, 2707,
	2709, 2730, 2732, 2738, 2740, 2741, 2743, 2747, 2751, 2751, 2770, 2770,
	2786, 2787, 2823, 2830, 2833, 2834, 2837, 2858, 2860, 2866, 2868, 2869,
	2871, 2875, 2879, 2879, 2910, 2911, 2913, 2915, 2931, 2931, 2949, 2949,
	2951, 2956, 2960, 2962, 2964, 2967, 2971, 2972, 2974, 2974, 2976, 2977,
	2981, 2982, 2986, 2988, 2992, 3003, 3026, 3026, 3079, 3086, 3088, 3090,
	3092, 3114, 3116, 3125, 3127, 3131, 3135, 3135, 3162, 3163, 3170, 3171,
	3207, 3214, 3216, 3218, 3220, 3242, 3244, 3253, 3255, 3259, 3263, 3263,
	3296, 3296, 3298, 3299, 3315, 3316, 3335, 3342, 3344, 3346, 3348, 3388,
	3391, 3391, 3408, 3408, 3426, 3427, 3452, 3457, 3463, 3480, 3484, 3507,
	3509, 3517, 3519, 3519, 3522, 3528, 3587, 3634, 3636, 3637, 3650, 3655,
	3715, 3716, 3718, 3718, 3721, 3722, 3724, 3724, 3727, 3727, 3734, 3737,
	3739, 3745, 3747, 3749, 3751, 3751, 3753, 3753, 3756, 3757, 3759, 3762,
	3764, 3765, 3775, 3775, 3778, 3782, 3806, 3809, 3842, 3842, 3906, 3913,
	3915, 3950, 3978, 3982, 4098, 4140, 4161, 4161, 4178, 4183, 4188, 4191,
	4195, 4195, 4199, 4200, 4208, 4210, 4215, 4227, 4240, 4240, 4306, 4348,
	4351, 4682, 4684, 4687, 4690, 4696, 4698, 4698, 4700, 4703, 4706, 4746,
	4748, 4751, 4754, 4786, 4788, 4791, 4794, 4800, 4802, 4802, 4804, 4807,
	4810, 4824, 4826, 4882, 4884, 4887, 4890, 4956, 4994, 5009, 5026, 5110,
	5123, 5742, 5745, 5761, 5763, 5788, 5794, 5868, 5890, 5902, 5904, 5907,
	5922, 5939, 5954, 5971, 5986, 5998, 6000, 6002, 6018, 6069, 6110, 6110,
	6178, 6212, 6214, 6265, 6274, 6314, 6316, 6316, 6322, 6391, 6402, 6430,
	6482, 6511, 6514, 6518, 6530, 6573, 6595, 6601, 6658, 6680, 6690, 6742,
	6919, 6965, 6983, 6989, 7045, 7074, 7088, 7089, 7100, 7143, 7170, 7205,
	7247, 7249, 7260, 7289, 7403, 7406, 7408, 7411, 7415, 7416, 8503, 8506,
	11570, 11625, 11650, 11672, 11682, 11688, 11690, 11696, 11698, 11704, 11706,
	11712, 11714, 11720, 11722, 11728, 11730, 11736, 11738, 11744, 12296, 12296,
	12350, 12350, 12355, 12440, 12449, 12449, 12451, 12540, 12545, 12545, 12551,
	12591, 12595, 12688, 12706, 12732, 12786, 12801, 13314, 13314, 19895, 19895,
	19970, 19970, 40910, 40910, 40962, 40982, 40984, 42126, 42194, 42233, 42242,
	42509, 42514, 42529, 42540, 42541, 42608, 42608, 42658, 42727, 43005, 43011,
	43013, 43015, 43017, 43020, 43022, 43044, 43074, 43125, 43140, 43189, 43252,
	43257, 43261, 43261, 43276, 43303, 43314, 43336, 43362, 43390, 43398, 43444,
	43522, 43562, 43586, 43588, 43590, 43597, 43618, 43633, 43635, 43640, 43644,
	43644, 43650, 43697, 43699, 43699, 43703, 43704, 43707, 43711, 43714, 43714,
	43716, 43716, 43741, 43742, 43746, 43756, 43764, 43764, 43779, 43784, 43787,
	43792, 43795, 43800, 43810, 43816, 43818, 43824, 43970, 44004, 44034, 44034,
	55205, 55205, 55218, 55240, 55245, 55293, 63746, 64111, 64114, 64219, 64287,
	64287, 64289, 64298, 64300, 64312, 64314, 64318, 64320, 64320, 64322, 64323,
	64325, 64326, 64328, 64435, 64469, 64831, 64850, 64913, 64916, 64969, 65010,
	65021, 65138, 65142, 65144, 65278, 65384, 65393, 65395, 65439, 65442, 65472,
	65476, 65481, 65484, 65489, 65492, 65497, 65500, 65502, 12, 2, 455, 455,
	458, 458, 461, 461, 500, 500, 8074, 8081, 8090, 8097, 8106, 8113, 8126,
	8126, 8142, 8142, 8190, 8190, 578, 2, 67, 92, 194, 216, 218, 224, 258,
	258, 260, 260, 262, 262, 264, 264, 266, 266, 268, 268, 270, 270, 272, 272,
	274, 274, 276, 276, 278, 278, 280, 280, 282, 282, 284, 284, 286, 286, 288,
	288, 290, 290, 292, 292, 294, 294, 296, 296, 298, 298, 300, 300, 302, 302,
	304, 304, 306, 306, 308, 308, 310, 310, 312, 312, 315, 315, 317, 317, 319,
	319, 321, 321, 323, 323, 325, 325, 327, 327, 329, 329, 332, 332, 334, 334,
	336, 336, 338, 338, 340, 340, 342, 342, 344, 344, 346, 346, 348, 348, 350,
	350, 352, 352, 354, 354, 356, 356, 358, 358, 360, 360, 362, 362, 364, 364,
	366, 366, 368, 368, 370, 370, 372, 372, 374, 374, 376, 376, 378, 379, 381,
	381, 383, 383, 387, 388, 390, 390, 392, 393, 395, 397, 400, 403, 405, 406,
	408, 410, 414, 415, 417, 418, 420, 420, 422, 422, 424, 425, 427, 427, 430,
	430, 432, 433, 435, 437, 439, 439, 441, 442, 446, 446, 454, 454, 457, 457,
	460, 460, 463, 463, 465, 465, 467, 467, 469, 469, 471, 471, 473, 473, 475,
	475, 477, 477, 480, 480, 482, 482, 484, 484, 486, 486, 488, 488, 490, 490,
	492, 492, 494, 494, 496, 496, 499, 499, 502, 502, 504, 506, 508, 508, 510,
	510, 512, 512, 514, 514, 516, 516, 518, 518, 520, 520, 522, 522, 524, 524,
	526, 526, 528, 528, 530, 530, 532, 532, 534, 534, 536, 536, 538, 538, 540,
	540, 542, 542, 544, 544, 546, 546, 548, 548, 550, 550, 552, 552, 554, 554,
	556, 556, 558, 558, 560, 560, 562, 562, 564, 564, 572, 573, 575, 576, 579,
	579, 581, 584, 586, 586, 588, 588, 590, 590, 592, 592, 882, 882, 884, 884,
	888, 888, 904, 904, 906, 908, 910, 910, 912, 913, 915, 931, 933, 941, 977,
	977, 980, 982, 986, 986, 988, 988, 990, 990, 992, 992, 994, 994, 996, 996,
	998, 998, 1000, 1000, 1002, 1002, 1004, 1004, 1006, 1006, 1008, 1008, 1014,
	1014, 1017, 1017, 1019, 1020, 1023, 1073, 1122, 1122, 1124, 1124, 1126,
	1126, 1128, 1128, 1130, 1130, 1132, 1132, 1134, 1134, 1136, 1136, 1138,
	1138, 1140, 1140, 1142, 1142, 1144, 1144, 1146, 1146, 1148, 1148, 1150,
	1150, 1152, 1152, 1154, 1154, 1164, 1164, 1166, 1166, 1168, 1168, 1170,
	1170, 1172, 1172, 1174, 1174, 1176, 1176, 1178, 1178, 1180, 1180, 1182,
	1182, 1184, 1184, 1186, 1186, 1188, 1188, 1190, 1190, 1192, 1192, 1194,
	1194, 1196, 1196, 1198, 1198, 1200, 1200, 1202, 1202, 1204, 1204, 1206,
	1206, 1208, 1208, 1210, 1210, 1212, 1212, 1214, 1214, 1216, 1216, 1218,
	1219, 1221, 1221, 1223, 1223, 1225, 1225, 1227, 1227, 1229, 1229, 1231,
	1231, 1234, 1234, 1236, 1236, 1238, 1238, 1240, 1240, 1242, 1242, 1244,
	1244, 1246, 1246, 1248, 1248, 1250, 1250, 1252, 1252, 1254, 1254, 1256,
	1256, 1258, 1258, 1260, 1260, 1262, 1262, 1264, 1264, 1266, 1266, 1268,
	1268, 1270, 1270, 1272, 1272, 1274, 1274, 1276, 1276, 1278, 1278, 1280,
	1280, 1282, 1282, 1284, 1284, 1286, 1286, 1288, 1288, 1290, 1290, 1292,
	1292, 1294, 1294, 1296, 1296, 1298, 1298, 1300, 1300, 1302, 1302, 1304,
	1304, 1306, 1306, 1308, 1308, 1310, 1310, 1312, 1312, 1314, 1314, 1316,
	1316, 1318, 1318, 1320, 1320, 1331, 1368, 4258, 4295, 4297, 4297, 4303,
	4303, 7682, 7682, 7684, 7684, 7686, 7686, 7688, 7688, 7690, 7690, 7692,
	7692, 7694, 7694, 7696, 7696, 7698, 7698, 7700, 7700, 7702, 7702, 7704,
	7704, 7706, 7706, 7708, 7708, 7710, 7710, 7712, 7712, 7714, 7714, 7716,
	7716, 7718, 7718, 7720, 7720, 7722, 7722, 7724, 7724, 7726, 7726, 7728,
	7728, 7730, 7730, 7732, 7732, 7734, 7734, 7736, 7736, 7738, 7738, 7740,
	7740, 7742, 7742, 7744, 7744, 7746, 7746, 7748, 7748, 7750, 7750, 7752,
	7752, 7754, 7754, 7756, 7756, 7758, 7758, 7760, 7760, 7762, 7762, 7764,
	7764, 7766, 7766, 7768, 7768, 7770, 7770, 7772, 7772, 7774, 7774, 7776,
	7776, 7778, 7778, 7780, 7780, 7782, 7782, 7784, 7784, 7786, 7786, 7788,
	7788, 7790, 7790, 7792, 7792, 7794, 7794, 7796, 7796, 7798, 7798, 7800,
	7800, 7802, 7802, 7804, 7804, 7806, 7806, 7808, 7808, 7810, 7810, 7812,
	7812, 7814, 7814, 7816, 7816, 7818, 7818, 7820, 7820, 7822, 7822, 7824,
	7824, 7826, 7826, 7828, 7828, 7830, 7830, 7840, 7840, 7842, 7842, 7844,
	7844, 7846, 7846, 7848, 7848, 7850, 7850, 7852, 7852, 7854, 7854, 7856,
	7856, 7858, 7858, 7860, 7860, 7862, 7862, 7864, 7864, 7866, 7866, 7868,
	7868, 7870, 7870, 7872, 7872, 7874, 7874, 7876, 7876, 7878, 7878, 7880,
	7880, 7882, 7882, 7884, 7884, 7886, 7886, 7888, 7888, 7890, 7890, 7892,
	7892, 7894, 7894, 7896, 7896, 7898, 7898, 7900, 7900, 7902, 7902, 7904,
	7904, 7906, 7906, 7908, 7908, 7910, 7910, 7912, 7912, 7914, 7914, 7916,
	7916, 7918, 7918, 7920, 7920, 7922, 7922, 7924, 7924, 7926, 7926, 7928,
	7928, 7930, 7930, 7932, 7932, 7934, 7934, 7936, 7936, 7946, 7953, 7962,
	7967, 7978, 7985, 7994, 8001, 8010, 8015, 8027, 8027, 8029, 8029, 8031,
	8031, 8033, 8033, 8042, 8049, 8122, 8125, 8138, 8141, 8154, 8157, 8170,
	8174, 8186, 8189, 8452, 8452, 8457, 8457, 8461, 8463, 8466, 8468, 8471,
	8471, 8475, 8479, 8486, 8486, 8488, 8488, 8490, 8490, 8492, 8495, 8498,
	8501, 8512, 8513, 8519, 8519, 8581, 8581, 11266, 11312, 11362, 11362, 11364,
	11366, 11369, 11369, 11371, 11371, 11373, 11373, 11375, 11378, 11380, 11380,
	11383, 11383, 11392, 11394, 11396, 11396, 11398, 11398, 11400, 11400, 11402,
	11402, 11404, 11404, 11406, 11406, 11408, 11408, 11410, 11410, 11412, 11412,
	11414, 11414, 11416, 11416, 11418, 11418, 11420, 11420, 11422, 11422, 11424,
	11424, 11426, 11426, 11428, 11428, 11430, 11430, 11432, 11432, 11434, 11434,
	11436, 11436, 11438, 11438, 11440, 11440, 11442, 11442, 11444, 11444, 11446,
	11446, 11448, 11448, 11450, 11450, 11452, 11452, 11454, 11454, 11456, 11456,
	11458, 11458, 11460, 11460, 11462, 11462, 11464, 11464, 11466, 11466, 11468,
	11468, 11470, 11470, 11472, 11472, 11474, 11474, 11476, 11476, 11478, 11478,
	11480, 11480, 11482, 11482, 11484, 11484, 11486, 11486, 11488, 11488, 11490,
	11490, 11492, 11492, 11501, 11501, 11503, 11503, 11508, 11508, 42562, 42562,
	42564, 42564, 42566, 42566, 42568, 42568, 42570, 42570, 42572, 42572, 42574,
	42574, 42576, 42576, 42578, 42578, 42580, 42580, 42582, 42582, 42584, 42584,
	42586, 42586, 42588, 42588, 42590, 42590, 42592, 42592, 42594, 42594, 42596,
	42596, 42598, 42598, 42600, 42600, 42602, 42602, 42604, 42604, 42606, 42606,
	42626, 42626, 42628, 42628, 42630, 42630, 42632, 42632, 42634, 42634, 42636,
	42636, 42638, 42638, 42640, 42640, 42642, 42642, 42644, 42644, 42646, 42646,
	42648, 42648, 42788, 42788, 42790, 42790, 42792, 42792, 42794, 42794, 42796,
	42796, 42798, 42798, 42800, 42800, 42804, 42804, 42806, 42806, 42808, 42808,
	42810, 42810, 42812, 42812, 42814, 42814, 42816, 42816, 42818, 42818, 42820,
	42820, 42822, 42822, 42824, 42824, 42826, 42826, 42828, 42828, 42830, 42830,
	42832, 42832, 42834, 42834, 42836, 42836, 42838, 42838, 42840, 42840, 42842,
	42842, 42844, 42844, 42846, 42846, 42848, 42848, 42850, 42850, 42852, 42852,
	42854, 42854, 42856, 42856, 42858, 42858, 42860, 42860, 42862, 42862, 42864,
	42864, 42875, 42875, 42877, 42877, 42879, 42880, 42882, 42882, 42884, 42884,
	42886, 42886, 42888, 42888, 42893, 42893, 42895, 42895, 42898, 42898, 42900,
	42900, 42914, 42914, 42916, 42916, 42918, 42918, 42920, 42920, 42922, 42922,
	42924, 42924, 65315, 65340, 37, 2, 50, 59, 1634, 1643, 1778, 1787, 1986,
	1995, 2408, 2417, 2536, 2545, 2664, 2673, 2792, 2801, 2920, 2929, 3048,
	3057, 3176, 3185, 3304, 3313, 3432, 3441, 3666, 3675, 3794, 3803, 3874,
	3883, 4162, 4171, 4242, 4251, 6114, 6123, 6162, 6171, 6472, 6481, 6610,
	6619, 6786, 6795, 6802, 6811, 6994, 7003, 7090, 7099, 7234, 7243, 7250,
	7259, 42530, 42539, 43218, 43227, 43266, 43275, 43474, 43483, 43602, 43611,
	44018, 44027, 65298, 65307, 9, 2, 5872, 5874, 8546, 8580, 8583, 8586, 12297,
	12297, 12323, 12331, 12346, 12348, 42728, 42737, 2, 30, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19, 3, 2,
	2, 2, 7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 25, 3, 2, 2, 2, 13, 27,
	3, 2, 2, 2, 15, 29, 3, 2, 2, 2, 17, 18, 9, 2, 2, 2, 18, 4, 3, 2, 2, 2,
	19, 20, 9, 3, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 9, 4, 2, 2, 22, 8, 3, 2,
	2, 2, 23, 24, 9, 5, 2, 2, 24, 10, 3, 2, 2, 2, 25, 26, 9, 6, 2, 2, 26, 12,
	3, 2, 2, 2, 27, 28, 9, 7, 2, 2, 28, 14, 3, 2, 2, 2, 29, 30, 9, 8, 2, 2,
	30, 16, 3, 2, 2, 2, 3, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "UNICODE_CLASS_LL", "UNICODE_CLASS_LM", "UNICODE_CLASS_LO", "UNICODE_CLASS_LT",
	"UNICODE_CLASS_LU", "UNICODE_CLASS_ND", "UNICODE_CLASS_NL",
}

var lexerRuleNames = []string{
	"UNICODE_CLASS_LL", "UNICODE_CLASS_LM", "UNICODE_CLASS_LO", "UNICODE_CLASS_LT",
	"UNICODE_CLASS_LU", "UNICODE_CLASS_ND", "UNICODE_CLASS_NL",
}

type UnicodeClasses struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewUnicodeClasses(input antlr.CharStream) *UnicodeClasses {

	l := new(UnicodeClasses)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "UnicodeClasses.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// UnicodeClasses tokens.
const (
	UnicodeClassesUNICODE_CLASS_LL = 1
	UnicodeClassesUNICODE_CLASS_LM = 2
	UnicodeClassesUNICODE_CLASS_LO = 3
	UnicodeClassesUNICODE_CLASS_LT = 4
	UnicodeClassesUNICODE_CLASS_LU = 5
	UnicodeClassesUNICODE_CLASS_ND = 6
	UnicodeClassesUNICODE_CLASS_NL = 7
)