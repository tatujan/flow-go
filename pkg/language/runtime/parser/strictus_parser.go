// Code generated from parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Strictus
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

var _ = strings.Builder{}

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 62, 640,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 3, 2, 7, 2, 134, 10, 2, 12, 2, 14, 2, 137, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 146, 10, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 152, 10, 4, 12, 4, 14, 4, 155, 11, 4, 3, 4, 5, 4, 158, 10, 4,
	3, 4, 3, 4, 5, 4, 162, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 167, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 174, 10, 6, 12, 6, 14, 6, 177, 11, 6, 3,
	6, 5, 6, 180, 10, 6, 3, 6, 7, 6, 183, 10, 6, 12, 6, 14, 6, 186, 11, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 194, 10, 7, 12, 7, 14, 7, 197,
	11, 7, 5, 7, 199, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 205, 10, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 215, 10, 10, 12, 10,
	14, 10, 218, 11, 10, 3, 10, 5, 10, 221, 10, 10, 3, 10, 7, 10, 224, 10,
	10, 12, 10, 14, 10, 227, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5,
	11, 234, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 244, 10, 12, 3, 12, 5, 12, 247, 10, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 255, 10, 13, 12, 13, 14, 13, 258, 11, 13, 5, 13,
	260, 10, 13, 3, 13, 3, 13, 3, 14, 5, 14, 265, 10, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 7, 15, 273, 10, 15, 12, 15, 14, 15, 276, 11, 15,
	3, 15, 7, 15, 279, 10, 15, 12, 15, 14, 15, 282, 11, 15, 3, 16, 3, 16, 5,
	16, 286, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 292, 10, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 299, 10, 18, 12, 18, 14, 18, 302, 11,
	18, 5, 18, 304, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 317, 10, 20, 3, 20, 5, 20, 320, 10,
	20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23, 338, 10, 23, 12, 23, 14,
	23, 341, 11, 23, 3, 24, 3, 24, 3, 24, 5, 24, 346, 10, 24, 3, 25, 3, 25,
	3, 25, 7, 25, 351, 10, 25, 12, 25, 14, 25, 354, 11, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 364, 10, 26, 3, 27, 3, 27,
	5, 27, 368, 10, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 5,
	30, 377, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 383, 10, 30, 5, 30,
	385, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 5,
	32, 395, 10, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 7, 33, 402, 10, 33,
	12, 33, 14, 33, 405, 11, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 418, 10, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 7, 36, 426, 10, 36, 12, 36, 14, 36, 429, 11, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 437, 10, 37, 12, 37, 14,
	37, 440, 11, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38,
	449, 10, 38, 12, 38, 14, 38, 452, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 7, 39, 461, 10, 39, 12, 39, 14, 39, 464, 11, 39, 3, 40,
	3, 40, 3, 40, 5, 40, 469, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 7, 41, 477, 10, 41, 12, 41, 14, 41, 480, 11, 41, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 7, 42, 489, 10, 42, 12, 42, 14, 42, 492, 11,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43, 501, 10, 43,
	12, 43, 14, 43, 504, 11, 43, 3, 44, 3, 44, 6, 44, 508, 10, 44, 13, 44,
	14, 44, 509, 3, 44, 3, 44, 5, 44, 514, 10, 44, 3, 45, 3, 45, 7, 45, 518,
	10, 45, 12, 45, 14, 45, 521, 11, 45, 3, 46, 3, 46, 5, 46, 525, 10, 46,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 5, 52, 543, 10, 52, 3, 52, 3, 52,
	3, 52, 3, 52, 3, 52, 3, 52, 5, 52, 551, 10, 52, 3, 53, 3, 53, 5, 53, 555,
	10, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56,
	3, 56, 3, 56, 7, 56, 568, 10, 56, 12, 56, 14, 56, 571, 11, 56, 5, 56, 573,
	10, 56, 3, 56, 3, 56, 3, 57, 3, 57, 5, 57, 579, 10, 57, 3, 57, 3, 57, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 589, 10, 58, 3, 59, 3, 59,
	3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 5, 62, 602,
	10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 7, 63, 608, 10, 63, 12, 63, 14, 63,
	611, 11, 63, 5, 63, 613, 10, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3,
	64, 7, 64, 621, 10, 64, 12, 64, 14, 64, 624, 11, 64, 5, 64, 626, 10, 64,
	3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 5,
	66, 638, 10, 66, 3, 66, 2, 9, 70, 72, 74, 76, 80, 82, 84, 67, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 2, 10, 3, 2, 44, 45, 3, 2,
	28, 29, 3, 2, 16, 17, 3, 2, 18, 21, 3, 2, 22, 23, 3, 2, 24, 26, 4, 2, 23,
	23, 27, 27, 3, 2, 49, 50, 2, 660, 2, 135, 3, 2, 2, 2, 4, 145, 3, 2, 2,
	2, 6, 147, 3, 2, 2, 2, 8, 166, 3, 2, 2, 2, 10, 168, 3, 2, 2, 2, 12, 198,
	3, 2, 2, 2, 14, 200, 3, 2, 2, 2, 16, 202, 3, 2, 2, 2, 18, 210, 3, 2, 2,
	2, 20, 230, 3, 2, 2, 2, 22, 237, 3, 2, 2, 2, 24, 250, 3, 2, 2, 2, 26, 264,
	3, 2, 2, 2, 28, 270, 3, 2, 2, 2, 30, 283, 3, 2, 2, 2, 32, 291, 3, 2, 2,
	2, 34, 293, 3, 2, 2, 2, 36, 310, 3, 2, 2, 2, 38, 314, 3, 2, 2, 2, 40, 324,
	3, 2, 2, 2, 42, 329, 3, 2, 2, 2, 44, 339, 3, 2, 2, 2, 46, 342, 3, 2, 2,
	2, 48, 352, 3, 2, 2, 2, 50, 363, 3, 2, 2, 2, 52, 365, 3, 2, 2, 2, 54, 369,
	3, 2, 2, 2, 56, 371, 3, 2, 2, 2, 58, 373, 3, 2, 2, 2, 60, 386, 3, 2, 2,
	2, 62, 390, 3, 2, 2, 2, 64, 399, 3, 2, 2, 2, 66, 409, 3, 2, 2, 2, 68, 411,
	3, 2, 2, 2, 70, 419, 3, 2, 2, 2, 72, 430, 3, 2, 2, 2, 74, 441, 3, 2, 2,
	2, 76, 453, 3, 2, 2, 2, 78, 465, 3, 2, 2, 2, 80, 470, 3, 2, 2, 2, 82, 481,
	3, 2, 2, 2, 84, 493, 3, 2, 2, 2, 86, 513, 3, 2, 2, 2, 88, 515, 3, 2, 2,
	2, 90, 524, 3, 2, 2, 2, 92, 526, 3, 2, 2, 2, 94, 528, 3, 2, 2, 2, 96, 530,
	3, 2, 2, 2, 98, 532, 3, 2, 2, 2, 100, 534, 3, 2, 2, 2, 102, 550, 3, 2,
	2, 2, 104, 554, 3, 2, 2, 2, 106, 556, 3, 2, 2, 2, 108, 559, 3, 2, 2, 2,
	110, 563, 3, 2, 2, 2, 112, 578, 3, 2, 2, 2, 114, 588, 3, 2, 2, 2, 116,
	590, 3, 2, 2, 2, 118, 592, 3, 2, 2, 2, 120, 594, 3, 2, 2, 2, 122, 601,
	3, 2, 2, 2, 124, 603, 3, 2, 2, 2, 126, 616, 3, 2, 2, 2, 128, 629, 3, 2,
	2, 2, 130, 637, 3, 2, 2, 2, 132, 134, 5, 4, 3, 2, 133, 132, 3, 2, 2, 2,
	134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	138, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 7, 2, 2, 3, 139, 3, 3,
	2, 2, 2, 140, 146, 5, 10, 6, 2, 141, 146, 5, 18, 10, 2, 142, 146, 5, 22,
	12, 2, 143, 146, 5, 62, 32, 2, 144, 146, 5, 6, 4, 2, 145, 140, 3, 2, 2,
	2, 145, 141, 3, 2, 2, 2, 145, 142, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145,
	144, 3, 2, 2, 2, 146, 5, 3, 2, 2, 2, 147, 157, 7, 3, 2, 2, 148, 153, 7,
	52, 2, 2, 149, 150, 7, 4, 2, 2, 150, 152, 7, 52, 2, 2, 151, 149, 3, 2,
	2, 2, 152, 155, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2,
	154, 156, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 156, 158, 7, 5, 2, 2, 157,
	148, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 162,
	5, 120, 61, 2, 160, 162, 7, 56, 2, 2, 161, 159, 3, 2, 2, 2, 161, 160, 3,
	2, 2, 2, 162, 7, 3, 2, 2, 2, 163, 167, 3, 2, 2, 2, 164, 167, 7, 39, 2,
	2, 165, 167, 7, 40, 2, 2, 166, 163, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166,
	165, 3, 2, 2, 2, 167, 9, 3, 2, 2, 2, 168, 169, 7, 34, 2, 2, 169, 170, 7,
	52, 2, 2, 170, 171, 5, 12, 7, 2, 171, 175, 7, 6, 2, 2, 172, 174, 5, 16,
	9, 2, 173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2,
	175, 176, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178,
	180, 5, 20, 11, 2, 179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 184,
	3, 2, 2, 2, 181, 183, 5, 22, 12, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3,
	2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2,
	2, 186, 184, 3, 2, 2, 2, 187, 188, 7, 7, 2, 2, 188, 11, 3, 2, 2, 2, 189,
	190, 7, 8, 2, 2, 190, 195, 7, 52, 2, 2, 191, 192, 7, 4, 2, 2, 192, 194,
	7, 52, 2, 2, 193, 191, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2,
	2, 2, 195, 196, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	198, 189, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 13, 3, 2, 2, 2, 200, 201,
	9, 2, 2, 2, 201, 15, 3, 2, 2, 2, 202, 204, 5, 8, 5, 2, 203, 205, 5, 14,
	8, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2,
	206, 207, 7, 52, 2, 2, 207, 208, 7, 8, 2, 2, 208, 209, 5, 28, 15, 2, 209,
	17, 3, 2, 2, 2, 210, 211, 7, 35, 2, 2, 211, 212, 7, 52, 2, 2, 212, 216,
	7, 6, 2, 2, 213, 215, 5, 16, 9, 2, 214, 213, 3, 2, 2, 2, 215, 218, 3, 2,
	2, 2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2,
	218, 216, 3, 2, 2, 2, 219, 221, 5, 20, 11, 2, 220, 219, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 225, 3, 2, 2, 2, 222, 224, 5, 22, 12, 2, 223, 222,
	3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2,
	2, 2, 226, 228, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 229, 7, 7, 2, 2,
	229, 19, 3, 2, 2, 2, 230, 231, 7, 52, 2, 2, 231, 233, 5, 24, 13, 2, 232,
	234, 5, 38, 20, 2, 233, 232, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235,
	3, 2, 2, 2, 235, 236, 6, 11, 2, 3, 236, 21, 3, 2, 2, 2, 237, 238, 5, 8,
	5, 2, 238, 239, 7, 36, 2, 2, 239, 240, 7, 52, 2, 2, 240, 243, 5, 24, 13,
	2, 241, 242, 7, 8, 2, 2, 242, 244, 5, 28, 15, 2, 243, 241, 3, 2, 2, 2,
	243, 244, 3, 2, 2, 2, 244, 246, 3, 2, 2, 2, 245, 247, 5, 38, 20, 2, 246,
	245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249,
	6, 12, 3, 3, 249, 23, 3, 2, 2, 2, 250, 259, 7, 31, 2, 2, 251, 256, 5, 26,
	14, 2, 252, 253, 7, 4, 2, 2, 253, 255, 5, 26, 14, 2, 254, 252, 3, 2, 2,
	2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 251, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 7, 32, 2, 2, 262, 25, 3, 2,
	2, 2, 263, 265, 7, 52, 2, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2,
	265, 266, 3, 2, 2, 2, 266, 267, 7, 52, 2, 2, 267, 268, 7, 8, 2, 2, 268,
	269, 5, 28, 15, 2, 269, 27, 3, 2, 2, 2, 270, 274, 5, 32, 17, 2, 271, 273,
	5, 30, 16, 2, 272, 271, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3,
	2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 280, 3, 2, 2, 2, 276, 274, 3, 2, 2,
	2, 277, 279, 9, 3, 2, 2, 278, 277, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280,
	278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 29, 3, 2, 2, 2, 282, 280, 3,
	2, 2, 2, 283, 285, 7, 9, 2, 2, 284, 286, 7, 53, 2, 2, 285, 284, 3, 2, 2,
	2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 7, 10, 2, 2, 288,
	31, 3, 2, 2, 2, 289, 292, 7, 52, 2, 2, 290, 292, 5, 34, 18, 2, 291, 289,
	3, 2, 2, 2, 291, 290, 3, 2, 2, 2, 292, 33, 3, 2, 2, 2, 293, 294, 7, 31,
	2, 2, 294, 303, 7, 31, 2, 2, 295, 300, 5, 28, 15, 2, 296, 297, 7, 4, 2,
	2, 297, 299, 5, 28, 15, 2, 298, 296, 3, 2, 2, 2, 299, 302, 3, 2, 2, 2,
	300, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302,
	300, 3, 2, 2, 2, 303, 295, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305,
	3, 2, 2, 2, 305, 306, 7, 32, 2, 2, 306, 307, 7, 8, 2, 2, 307, 308, 5, 28,
	15, 2, 308, 309, 7, 32, 2, 2, 309, 35, 3, 2, 2, 2, 310, 311, 7, 6, 2, 2,
	311, 312, 5, 48, 25, 2, 312, 313, 7, 7, 2, 2, 313, 37, 3, 2, 2, 2, 314,
	316, 7, 6, 2, 2, 315, 317, 5, 40, 21, 2, 316, 315, 3, 2, 2, 2, 316, 317,
	3, 2, 2, 2, 317, 319, 3, 2, 2, 2, 318, 320, 5, 42, 22, 2, 319, 318, 3,
	2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 322, 5, 48, 25,
	2, 322, 323, 7, 7, 2, 2, 323, 39, 3, 2, 2, 2, 324, 325, 7, 37, 2, 2, 325,
	326, 7, 6, 2, 2, 326, 327, 5, 44, 23, 2, 327, 328, 7, 7, 2, 2, 328, 41,
	3, 2, 2, 2, 329, 330, 7, 38, 2, 2, 330, 331, 7, 6, 2, 2, 331, 332, 5, 44,
	23, 2, 332, 333, 7, 7, 2, 2, 333, 43, 3, 2, 2, 2, 334, 335, 5, 46, 24,
	2, 335, 336, 5, 130, 66, 2, 336, 338, 3, 2, 2, 2, 337, 334, 3, 2, 2, 2,
	338, 341, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340,
	45, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 342, 345, 5, 66, 34, 2, 343, 344,
	7, 8, 2, 2, 344, 346, 5, 66, 34, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3,
	2, 2, 2, 346, 47, 3, 2, 2, 2, 347, 348, 5, 50, 26, 2, 348, 349, 5, 130,
	66, 2, 349, 351, 3, 2, 2, 2, 350, 347, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2,
	352, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 49, 3, 2, 2, 2, 354, 352,
	3, 2, 2, 2, 355, 364, 5, 52, 27, 2, 356, 364, 5, 54, 28, 2, 357, 364, 5,
	56, 29, 2, 358, 364, 5, 58, 30, 2, 359, 364, 5, 60, 31, 2, 360, 364, 5,
	4, 3, 2, 361, 364, 5, 64, 33, 2, 362, 364, 5, 66, 34, 2, 363, 355, 3, 2,
	2, 2, 363, 356, 3, 2, 2, 2, 363, 357, 3, 2, 2, 2, 363, 358, 3, 2, 2, 2,
	363, 359, 3, 2, 2, 2, 363, 360, 3, 2, 2, 2, 363, 361, 3, 2, 2, 2, 363,
	362, 3, 2, 2, 2, 364, 51, 3, 2, 2, 2, 365, 367, 7, 41, 2, 2, 366, 368,
	5, 66, 34, 2, 367, 366, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 53, 3, 2,
	2, 2, 369, 370, 7, 42, 2, 2, 370, 55, 3, 2, 2, 2, 371, 372, 7, 43, 2, 2,
	372, 57, 3, 2, 2, 2, 373, 376, 7, 46, 2, 2, 374, 377, 5, 66, 34, 2, 375,
	377, 5, 62, 32, 2, 376, 374, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 378,
	3, 2, 2, 2, 378, 384, 5, 36, 19, 2, 379, 382, 7, 47, 2, 2, 380, 383, 5,
	58, 30, 2, 381, 383, 5, 36, 19, 2, 382, 380, 3, 2, 2, 2, 382, 381, 3, 2,
	2, 2, 383, 385, 3, 2, 2, 2, 384, 379, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2,
	385, 59, 3, 2, 2, 2, 386, 387, 7, 48, 2, 2, 387, 388, 5, 66, 34, 2, 388,
	389, 5, 36, 19, 2, 389, 61, 3, 2, 2, 2, 390, 391, 5, 14, 8, 2, 391, 394,
	7, 52, 2, 2, 392, 393, 7, 8, 2, 2, 393, 395, 5, 28, 15, 2, 394, 392, 3,
	2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 397, 7, 11, 2,
	2, 397, 398, 5, 66, 34, 2, 398, 63, 3, 2, 2, 2, 399, 403, 7, 52, 2, 2,
	400, 402, 5, 104, 53, 2, 401, 400, 3, 2, 2, 2, 402, 405, 3, 2, 2, 2, 403,
	401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406, 3, 2, 2, 2, 405, 403,
	3, 2, 2, 2, 406, 407, 7, 11, 2, 2, 407, 408, 5, 66, 34, 2, 408, 65, 3,
	2, 2, 2, 409, 410, 5, 68, 35, 2, 410, 67, 3, 2, 2, 2, 411, 417, 5, 70,
	36, 2, 412, 413, 7, 28, 2, 2, 413, 414, 5, 66, 34, 2, 414, 415, 7, 8, 2,
	2, 415, 416, 5, 66, 34, 2, 416, 418, 3, 2, 2, 2, 417, 412, 3, 2, 2, 2,
	417, 418, 3, 2, 2, 2, 418, 69, 3, 2, 2, 2, 419, 420, 8, 36, 1, 2, 420,
	421, 5, 72, 37, 2, 421, 427, 3, 2, 2, 2, 422, 423, 12, 3, 2, 2, 423, 424,
	7, 12, 2, 2, 424, 426, 5, 72, 37, 2, 425, 422, 3, 2, 2, 2, 426, 429, 3,
	2, 2, 2, 427, 425, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 71, 3, 2, 2,
	2, 429, 427, 3, 2, 2, 2, 430, 431, 8, 37, 1, 2, 431, 432, 5, 74, 38, 2,
	432, 438, 3, 2, 2, 2, 433, 434, 12, 3, 2, 2, 434, 435, 7, 13, 2, 2, 435,
	437, 5, 74, 38, 2, 436, 433, 3, 2, 2, 2, 437, 440, 3, 2, 2, 2, 438, 436,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 73, 3, 2, 2, 2, 440, 438, 3, 2,
	2, 2, 441, 442, 8, 38, 1, 2, 442, 443, 5, 76, 39, 2, 443, 450, 3, 2, 2,
	2, 444, 445, 12, 3, 2, 2, 445, 446, 5, 92, 47, 2, 446, 447, 5, 76, 39,
	2, 447, 449, 3, 2, 2, 2, 448, 444, 3, 2, 2, 2, 449, 452, 3, 2, 2, 2, 450,
	448, 3, 2, 2, 2, 450, 451, 3, 2, 2, 2, 451, 75, 3, 2, 2, 2, 452, 450, 3,
	2, 2, 2, 453, 454, 8, 39, 1, 2, 454, 455, 5, 78, 40, 2, 455, 462, 3, 2,
	2, 2, 456, 457, 12, 3, 2, 2, 457, 458, 5, 94, 48, 2, 458, 459, 5, 78, 40,
	2, 459, 461, 3, 2, 2, 2, 460, 456, 3, 2, 2, 2, 461, 464, 3, 2, 2, 2, 462,
	460, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 77, 3, 2, 2, 2, 464, 462, 3,
	2, 2, 2, 465, 468, 5, 80, 41, 2, 466, 467, 7, 29, 2, 2, 467, 469, 5, 78,
	40, 2, 468, 466, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 79, 3, 2, 2, 2,
	470, 471, 8, 41, 1, 2, 471, 472, 5, 82, 42, 2, 472, 478, 3, 2, 2, 2, 473,
	474, 12, 3, 2, 2, 474, 475, 7, 30, 2, 2, 475, 477, 5, 28, 15, 2, 476, 473,
	3, 2, 2, 2, 477, 480, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 478, 479, 3, 2,
	2, 2, 479, 81, 3, 2, 2, 2, 480, 478, 3, 2, 2, 2, 481, 482, 8, 42, 1, 2,
	482, 483, 5, 84, 43, 2, 483, 490, 3, 2, 2, 2, 484, 485, 12, 3, 2, 2, 485,
	486, 5, 96, 49, 2, 486, 487, 5, 84, 43, 2, 487, 489, 3, 2, 2, 2, 488, 484,
	3, 2, 2, 2, 489, 492, 3, 2, 2, 2, 490, 488, 3, 2, 2, 2, 490, 491, 3, 2,
	2, 2, 491, 83, 3, 2, 2, 2, 492, 490, 3, 2, 2, 2, 493, 494, 8, 43, 1, 2,
	494, 495, 5, 86, 44, 2, 495, 502, 3, 2, 2, 2, 496, 497, 12, 3, 2, 2, 497,
	498, 5, 98, 50, 2, 498, 499, 5, 86, 44, 2, 499, 501, 3, 2, 2, 2, 500, 496,
	3, 2, 2, 2, 501, 504, 3, 2, 2, 2, 502, 500, 3, 2, 2, 2, 502, 503, 3, 2,
	2, 2, 503, 85, 3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 505, 514, 5, 88, 45, 2,
	506, 508, 5, 100, 51, 2, 507, 506, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509,
	507, 3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510, 511, 3, 2, 2, 2, 511, 512,
	5, 86, 44, 2, 512, 514, 3, 2, 2, 2, 513, 505, 3, 2, 2, 2, 513, 507, 3,
	2, 2, 2, 514, 87, 3, 2, 2, 2, 515, 519, 5, 102, 52, 2, 516, 518, 5, 90,
	46, 2, 517, 516, 3, 2, 2, 2, 518, 521, 3, 2, 2, 2, 519, 517, 3, 2, 2, 2,
	519, 520, 3, 2, 2, 2, 520, 89, 3, 2, 2, 2, 521, 519, 3, 2, 2, 2, 522, 525,
	5, 104, 53, 2, 523, 525, 5, 110, 56, 2, 524, 522, 3, 2, 2, 2, 524, 523,
	3, 2, 2, 2, 525, 91, 3, 2, 2, 2, 526, 527, 9, 4, 2, 2, 527, 93, 3, 2, 2,
	2, 528, 529, 9, 5, 2, 2, 529, 95, 3, 2, 2, 2, 530, 531, 9, 6, 2, 2, 531,
	97, 3, 2, 2, 2, 532, 533, 9, 7, 2, 2, 533, 99, 3, 2, 2, 2, 534, 535, 9,
	8, 2, 2, 535, 101, 3, 2, 2, 2, 536, 551, 7, 52, 2, 2, 537, 551, 5, 114,
	58, 2, 538, 539, 7, 36, 2, 2, 539, 542, 5, 24, 13, 2, 540, 541, 7, 8, 2,
	2, 541, 543, 5, 28, 15, 2, 542, 540, 3, 2, 2, 2, 542, 543, 3, 2, 2, 2,
	543, 544, 3, 2, 2, 2, 544, 545, 5, 38, 20, 2, 545, 551, 3, 2, 2, 2, 546,
	547, 7, 31, 2, 2, 547, 548, 5, 66, 34, 2, 548, 549, 7, 32, 2, 2, 549, 551,
	3, 2, 2, 2, 550, 536, 3, 2, 2, 2, 550, 537, 3, 2, 2, 2, 550, 538, 3, 2,
	2, 2, 550, 546, 3, 2, 2, 2, 551, 103, 3, 2, 2, 2, 552, 555, 5, 106, 54,
	2, 553, 555, 5, 108, 55, 2, 554, 552, 3, 2, 2, 2, 554, 553, 3, 2, 2, 2,
	555, 105, 3, 2, 2, 2, 556, 557, 7, 14, 2, 2, 557, 558, 7, 52, 2, 2, 558,
	107, 3, 2, 2, 2, 559, 560, 7, 9, 2, 2, 560, 561, 5, 66, 34, 2, 561, 562,
	7, 10, 2, 2, 562, 109, 3, 2, 2, 2, 563, 572, 7, 31, 2, 2, 564, 569, 5,
	112, 57, 2, 565, 566, 7, 4, 2, 2, 566, 568, 5, 112, 57, 2, 567, 565, 3,
	2, 2, 2, 568, 571, 3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 569, 570, 3, 2, 2,
	2, 570, 573, 3, 2, 2, 2, 571, 569, 3, 2, 2, 2, 572, 564, 3, 2, 2, 2, 572,
	573, 3, 2, 2, 2, 573, 574, 3, 2, 2, 2, 574, 575, 7, 32, 2, 2, 575, 111,
	3, 2, 2, 2, 576, 577, 7, 52, 2, 2, 577, 579, 7, 8, 2, 2, 578, 576, 3, 2,
	2, 2, 578, 579, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2, 580, 581, 5, 66, 34,
	2, 581, 113, 3, 2, 2, 2, 582, 589, 5, 122, 62, 2, 583, 589, 5, 116, 59,
	2, 584, 589, 5, 124, 63, 2, 585, 589, 5, 126, 64, 2, 586, 589, 5, 120,
	61, 2, 587, 589, 5, 118, 60, 2, 588, 582, 3, 2, 2, 2, 588, 583, 3, 2, 2,
	2, 588, 584, 3, 2, 2, 2, 588, 585, 3, 2, 2, 2, 588, 586, 3, 2, 2, 2, 588,
	587, 3, 2, 2, 2, 589, 115, 3, 2, 2, 2, 590, 591, 9, 9, 2, 2, 591, 117,
	3, 2, 2, 2, 592, 593, 7, 51, 2, 2, 593, 119, 3, 2, 2, 2, 594, 595, 7, 58,
	2, 2, 595, 121, 3, 2, 2, 2, 596, 602, 7, 53, 2, 2, 597, 602, 7, 54, 2,
	2, 598, 602, 7, 55, 2, 2, 599, 602, 7, 56, 2, 2, 600, 602, 7, 57, 2, 2,
	601, 596, 3, 2, 2, 2, 601, 597, 3, 2, 2, 2, 601, 598, 3, 2, 2, 2, 601,
	599, 3, 2, 2, 2, 601, 600, 3, 2, 2, 2, 602, 123, 3, 2, 2, 2, 603, 612,
	7, 9, 2, 2, 604, 609, 5, 66, 34, 2, 605, 606, 7, 4, 2, 2, 606, 608, 5,
	66, 34, 2, 607, 605, 3, 2, 2, 2, 608, 611, 3, 2, 2, 2, 609, 607, 3, 2,
	2, 2, 609, 610, 3, 2, 2, 2, 610, 613, 3, 2, 2, 2, 611, 609, 3, 2, 2, 2,
	612, 604, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614,
	615, 7, 10, 2, 2, 615, 125, 3, 2, 2, 2, 616, 625, 7, 6, 2, 2, 617, 622,
	5, 128, 65, 2, 618, 619, 7, 4, 2, 2, 619, 621, 5, 128, 65, 2, 620, 618,
	3, 2, 2, 2, 621, 624, 3, 2, 2, 2, 622, 620, 3, 2, 2, 2, 622, 623, 3, 2,
	2, 2, 623, 626, 3, 2, 2, 2, 624, 622, 3, 2, 2, 2, 625, 617, 3, 2, 2, 2,
	625, 626, 3, 2, 2, 2, 626, 627, 3, 2, 2, 2, 627, 628, 7, 7, 2, 2, 628,
	127, 3, 2, 2, 2, 629, 630, 5, 66, 34, 2, 630, 631, 7, 8, 2, 2, 631, 632,
	5, 66, 34, 2, 632, 129, 3, 2, 2, 2, 633, 638, 7, 15, 2, 2, 634, 638, 7,
	2, 2, 3, 635, 638, 6, 66, 11, 2, 636, 638, 6, 66, 12, 2, 637, 633, 3, 2,
	2, 2, 637, 634, 3, 2, 2, 2, 637, 635, 3, 2, 2, 2, 637, 636, 3, 2, 2, 2,
	638, 131, 3, 2, 2, 2, 67, 135, 145, 153, 157, 161, 166, 175, 179, 184,
	195, 198, 204, 216, 220, 225, 233, 243, 246, 256, 259, 264, 274, 280, 285,
	291, 300, 303, 316, 319, 339, 345, 352, 363, 367, 376, 382, 384, 394, 403,
	417, 427, 438, 450, 462, 468, 478, 490, 502, 509, 513, 519, 524, 542, 550,
	554, 569, 572, 578, 588, 601, 609, 612, 622, 625, 637,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'import'", "','", "'from'", "'{'", "'}'", "':'", "'['", "']'", "'='",
	"'||'", "'&&'", "'.'", "';'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
	"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'?'", "'??'", "'as?'", "'('",
	"')'", "'transaction'", "'struct'", "'interface'", "'fun'", "'pre'", "'post'",
	"'pub'", "'pub(set)'", "'return'", "'break'", "'continue'", "'let'", "'var'",
	"'if'", "'else'", "'while'", "'true'", "'false'", "'nil'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal",
	"Less", "Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul",
	"Div", "Mod", "Negate", "Optional", "NilCoalescing", "FailableDowncasting",
	"OpenParen", "CloseParen", "Transaction", "Struct", "Interface", "Fun",
	"Pre", "Post", "Pub", "PubSet", "Return", "Break", "Continue", "Let", "Var",
	"If", "Else", "While", "True", "False", "Nil", "Identifier", "DecimalLiteral",
	"BinaryLiteral", "OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral",
	"StringLiteral", "WS", "Terminator", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "declaration", "importDeclaration", "access", "structureDeclaration",
	"conformances", "variableKind", "field", "interfaceDeclaration", "initializer",
	"functionDeclaration", "parameterList", "parameter", "fullType", "typeDimension",
	"baseType", "functionType", "block", "functionBlock", "preConditions",
	"postConditions", "conditions", "condition", "statements", "statement",
	"returnStatement", "breakStatement", "continueStatement", "ifStatement",
	"whileStatement", "variableDeclaration", "assignment", "expression", "conditionalExpression",
	"orExpression", "andExpression", "equalityExpression", "relationalExpression",
	"nilCoalescingExpression", "failableDowncastingExpression", "additiveExpression",
	"multiplicativeExpression", "unaryExpression", "primaryExpression", "primaryExpressionSuffix",
	"equalityOp", "relationalOp", "additiveOp", "multiplicativeOp", "unaryOp",
	"primaryExpressionStart", "expressionAccess", "memberAccess", "bracketExpression",
	"invocation", "argument", "literal", "booleanLiteral", "nilLiteral", "stringLiteral",
	"integerLiteral", "arrayLiteral", "dictionaryLiteral", "dictionaryEntry",
	"eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type StrictusParser struct {
	*antlr.BaseParser
}

func NewStrictusParser(input antlr.TokenStream) *StrictusParser {
	this := new(StrictusParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Strictus.g4"

	return this
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *StrictusParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return true
	}

	if ahead.GetTokenType() == StrictusParserTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == StrictusParserWS {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == StrictusParserBlockComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == StrictusParserTerminator)
}

// StrictusParser tokens.
const (
	StrictusParserEOF                  = antlr.TokenEOF
	StrictusParserT__0                 = 1
	StrictusParserT__1                 = 2
	StrictusParserT__2                 = 3
	StrictusParserT__3                 = 4
	StrictusParserT__4                 = 5
	StrictusParserT__5                 = 6
	StrictusParserT__6                 = 7
	StrictusParserT__7                 = 8
	StrictusParserT__8                 = 9
	StrictusParserT__9                 = 10
	StrictusParserT__10                = 11
	StrictusParserT__11                = 12
	StrictusParserT__12                = 13
	StrictusParserEqual                = 14
	StrictusParserUnequal              = 15
	StrictusParserLess                 = 16
	StrictusParserGreater              = 17
	StrictusParserLessEqual            = 18
	StrictusParserGreaterEqual         = 19
	StrictusParserPlus                 = 20
	StrictusParserMinus                = 21
	StrictusParserMul                  = 22
	StrictusParserDiv                  = 23
	StrictusParserMod                  = 24
	StrictusParserNegate               = 25
	StrictusParserOptional             = 26
	StrictusParserNilCoalescing        = 27
	StrictusParserFailableDowncasting  = 28
	StrictusParserOpenParen            = 29
	StrictusParserCloseParen           = 30
	StrictusParserTransaction          = 31
	StrictusParserStruct               = 32
	StrictusParserInterface            = 33
	StrictusParserFun                  = 34
	StrictusParserPre                  = 35
	StrictusParserPost                 = 36
	StrictusParserPub                  = 37
	StrictusParserPubSet               = 38
	StrictusParserReturn               = 39
	StrictusParserBreak                = 40
	StrictusParserContinue             = 41
	StrictusParserLet                  = 42
	StrictusParserVar                  = 43
	StrictusParserIf                   = 44
	StrictusParserElse                 = 45
	StrictusParserWhile                = 46
	StrictusParserTrue                 = 47
	StrictusParserFalse                = 48
	StrictusParserNil                  = 49
	StrictusParserIdentifier           = 50
	StrictusParserDecimalLiteral       = 51
	StrictusParserBinaryLiteral        = 52
	StrictusParserOctalLiteral         = 53
	StrictusParserHexadecimalLiteral   = 54
	StrictusParserInvalidNumberLiteral = 55
	StrictusParserStringLiteral        = 56
	StrictusParserWS                   = 57
	StrictusParserTerminator           = 58
	StrictusParserBlockComment         = 59
	StrictusParserLineComment          = 60
)

// StrictusParser rules.
const (
	StrictusParserRULE_program                       = 0
	StrictusParserRULE_declaration                   = 1
	StrictusParserRULE_importDeclaration             = 2
	StrictusParserRULE_access                        = 3
	StrictusParserRULE_structureDeclaration          = 4
	StrictusParserRULE_conformances                  = 5
	StrictusParserRULE_variableKind                  = 6
	StrictusParserRULE_field                         = 7
	StrictusParserRULE_interfaceDeclaration          = 8
	StrictusParserRULE_initializer                   = 9
	StrictusParserRULE_functionDeclaration           = 10
	StrictusParserRULE_parameterList                 = 11
	StrictusParserRULE_parameter                     = 12
	StrictusParserRULE_fullType                      = 13
	StrictusParserRULE_typeDimension                 = 14
	StrictusParserRULE_baseType                      = 15
	StrictusParserRULE_functionType                  = 16
	StrictusParserRULE_block                         = 17
	StrictusParserRULE_functionBlock                 = 18
	StrictusParserRULE_preConditions                 = 19
	StrictusParserRULE_postConditions                = 20
	StrictusParserRULE_conditions                    = 21
	StrictusParserRULE_condition                     = 22
	StrictusParserRULE_statements                    = 23
	StrictusParserRULE_statement                     = 24
	StrictusParserRULE_returnStatement               = 25
	StrictusParserRULE_breakStatement                = 26
	StrictusParserRULE_continueStatement             = 27
	StrictusParserRULE_ifStatement                   = 28
	StrictusParserRULE_whileStatement                = 29
	StrictusParserRULE_variableDeclaration           = 30
	StrictusParserRULE_assignment                    = 31
	StrictusParserRULE_expression                    = 32
	StrictusParserRULE_conditionalExpression         = 33
	StrictusParserRULE_orExpression                  = 34
	StrictusParserRULE_andExpression                 = 35
	StrictusParserRULE_equalityExpression            = 36
	StrictusParserRULE_relationalExpression          = 37
	StrictusParserRULE_nilCoalescingExpression       = 38
	StrictusParserRULE_failableDowncastingExpression = 39
	StrictusParserRULE_additiveExpression            = 40
	StrictusParserRULE_multiplicativeExpression      = 41
	StrictusParserRULE_unaryExpression               = 42
	StrictusParserRULE_primaryExpression             = 43
	StrictusParserRULE_primaryExpressionSuffix       = 44
	StrictusParserRULE_equalityOp                    = 45
	StrictusParserRULE_relationalOp                  = 46
	StrictusParserRULE_additiveOp                    = 47
	StrictusParserRULE_multiplicativeOp              = 48
	StrictusParserRULE_unaryOp                       = 49
	StrictusParserRULE_primaryExpressionStart        = 50
	StrictusParserRULE_expressionAccess              = 51
	StrictusParserRULE_memberAccess                  = 52
	StrictusParserRULE_bracketExpression             = 53
	StrictusParserRULE_invocation                    = 54
	StrictusParserRULE_argument                      = 55
	StrictusParserRULE_literal                       = 56
	StrictusParserRULE_booleanLiteral                = 57
	StrictusParserRULE_nilLiteral                    = 58
	StrictusParserRULE_stringLiteral                 = 59
	StrictusParserRULE_integerLiteral                = 60
	StrictusParserRULE_arrayLiteral                  = 61
	StrictusParserRULE_dictionaryLiteral             = 62
	StrictusParserRULE_dictionaryEntry               = 63
	StrictusParserRULE_eos                           = 64
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StrictusParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StrictusParserT__0 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(StrictusParserStruct-32))|(1<<(StrictusParserInterface-32))|(1<<(StrictusParserFun-32))|(1<<(StrictusParserPub-32))|(1<<(StrictusParserPubSet-32))|(1<<(StrictusParserLet-32))|(1<<(StrictusParserVar-32)))) != 0) {
		{
			p.SetState(130)
			p.Declaration()
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(StrictusParserEOF)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) StructureDeclaration() IStructureDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructureDeclarationContext)
}

func (s *DeclarationContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StrictusParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserStruct:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.StructureDeclaration()
		}

	case StrictusParserInterface:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.InterfaceDeclaration()
		}

	case StrictusParserFun, StrictusParserPub, StrictusParserPubSet:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.FunctionDeclaration(true)
		}

	case StrictusParserLet, StrictusParserVar:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.VariableDeclaration()
		}

	case StrictusParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.ImportDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ImportDeclarationContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserHexadecimalLiteral, 0)
}

func (s *ImportDeclarationContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserIdentifier)
}

func (s *ImportDeclarationContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, i)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, StrictusParserRULE_importDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(StrictusParserT__0)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserIdentifier {
		{
			p.SetState(146)
			p.Match(StrictusParserIdentifier)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(147)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(148)
				p.Match(StrictusParserIdentifier)
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.Match(StrictusParserT__2)
		}

	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserStringLiteral:
		{
			p.SetState(157)
			p.StringLiteral()
		}

	case StrictusParserHexadecimalLiteral:
		{
			p.SetState(158)
			p.Match(StrictusParserHexadecimalLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAccessContext is an interface to support dynamic dispatch.
type IAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessContext differentiates from other interfaces.
	IsAccessContext()
}

type AccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessContext() *AccessContext {
	var p = new(AccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_access
	return p
}

func (*AccessContext) IsAccessContext() {}

func NewAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessContext {
	var p = new(AccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_access

	return p
}

func (s *AccessContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessContext) Pub() antlr.TerminalNode {
	return s.GetToken(StrictusParserPub, 0)
}

func (s *AccessContext) PubSet() antlr.TerminalNode {
	return s.GetToken(StrictusParserPubSet, 0)
}

func (s *AccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAccess(s)
	}
}

func (s *AccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAccess(s)
	}
}

func (s *AccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Access() (localctx IAccessContext) {
	localctx = NewAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, StrictusParserRULE_access)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserFun, StrictusParserLet, StrictusParserVar, StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)

	case StrictusParserPub:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(StrictusParserPub)
		}

	case StrictusParserPubSet:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Match(StrictusParserPubSet)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStructureDeclarationContext is an interface to support dynamic dispatch.
type IStructureDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureDeclarationContext differentiates from other interfaces.
	IsStructureDeclarationContext()
}

type StructureDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureDeclarationContext() *StructureDeclarationContext {
	var p = new(StructureDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_structureDeclaration
	return p
}

func (*StructureDeclarationContext) IsStructureDeclarationContext() {}

func NewStructureDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureDeclarationContext {
	var p = new(StructureDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_structureDeclaration

	return p
}

func (s *StructureDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureDeclarationContext) Struct() antlr.TerminalNode {
	return s.GetToken(StrictusParserStruct, 0)
}

func (s *StructureDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *StructureDeclarationContext) Conformances() IConformancesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConformancesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConformancesContext)
}

func (s *StructureDeclarationContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *StructureDeclarationContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StructureDeclarationContext) Initializer() IInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *StructureDeclarationContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem())
	var tst = make([]IFunctionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclarationContext)
		}
	}

	return tst
}

func (s *StructureDeclarationContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StructureDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructureDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStructureDeclaration(s)
	}
}

func (s *StructureDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStructureDeclaration(s)
	}
}

func (s *StructureDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStructureDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) StructureDeclaration() (localctx IStructureDeclarationContext) {
	localctx = NewStructureDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, StrictusParserRULE_structureDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(StrictusParserStruct)
	}
	{
		p.SetState(167)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(168)
		p.Conformances()
	}
	{
		p.SetState(169)
		p.Match(StrictusParserT__3)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.Field()
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserIdentifier {
		{
			p.SetState(176)
			p.Initializer(true)
		}

	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserPub-34))|(1<<(StrictusParserPubSet-34)))) != 0 {
		{
			p.SetState(179)
			p.FunctionDeclaration(true)
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IConformancesContext is an interface to support dynamic dispatch.
type IConformancesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConformancesContext differentiates from other interfaces.
	IsConformancesContext()
}

type ConformancesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConformancesContext() *ConformancesContext {
	var p = new(ConformancesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conformances
	return p
}

func (*ConformancesContext) IsConformancesContext() {}

func NewConformancesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConformancesContext {
	var p = new(ConformancesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conformances

	return p
}

func (s *ConformancesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConformancesContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserIdentifier)
}

func (s *ConformancesContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, i)
}

func (s *ConformancesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConformancesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConformancesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConformances(s)
	}
}

func (s *ConformancesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConformances(s)
	}
}

func (s *ConformancesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConformances(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Conformances() (localctx IConformancesContext) {
	localctx = NewConformancesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, StrictusParserRULE_conformances)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__5 {
		{
			p.SetState(187)
			p.Match(StrictusParserT__5)
		}
		{
			p.SetState(188)
			p.Match(StrictusParserIdentifier)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(189)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(190)
				p.Match(StrictusParserIdentifier)
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IVariableKindContext is an interface to support dynamic dispatch.
type IVariableKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableKindContext differentiates from other interfaces.
	IsVariableKindContext()
}

type VariableKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableKindContext() *VariableKindContext {
	var p = new(VariableKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableKind
	return p
}

func (*VariableKindContext) IsVariableKindContext() {}

func NewVariableKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableKindContext {
	var p = new(VariableKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableKind

	return p
}

func (s *VariableKindContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableKindContext) Let() antlr.TerminalNode {
	return s.GetToken(StrictusParserLet, 0)
}

func (s *VariableKindContext) Var() antlr.TerminalNode {
	return s.GetToken(StrictusParserVar, 0)
}

func (s *VariableKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableKind(s)
	}
}

func (s *VariableKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableKind(s)
	}
}

func (s *VariableKindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableKind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableKind() (localctx IVariableKindContext) {
	localctx = NewVariableKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, StrictusParserRULE_variableKind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserLet || _la == StrictusParserVar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FieldContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *FieldContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *FieldContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, StrictusParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Access()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserLet || _la == StrictusParserVar {
		{
			p.SetState(201)
			p.VariableKind()
		}

	}
	{
		p.SetState(204)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(205)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(206)
		p.FullType()
	}

	return localctx
}

// IInterfaceDeclarationContext is an interface to support dynamic dispatch.
type IInterfaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceDeclarationContext differentiates from other interfaces.
	IsInterfaceDeclarationContext()
}

type InterfaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclarationContext() *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_interfaceDeclaration
	return p
}

func (*InterfaceDeclarationContext) IsInterfaceDeclarationContext() {}

func NewInterfaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_interfaceDeclaration

	return p
}

func (s *InterfaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclarationContext) Interface() antlr.TerminalNode {
	return s.GetToken(StrictusParserInterface, 0)
}

func (s *InterfaceDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *InterfaceDeclarationContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *InterfaceDeclarationContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *InterfaceDeclarationContext) Initializer() IInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitializerContext)
}

func (s *InterfaceDeclarationContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem())
	var tst = make([]IFunctionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclarationContext)
		}
	}

	return tst
}

func (s *InterfaceDeclarationContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *InterfaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInterfaceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) InterfaceDeclaration() (localctx IInterfaceDeclarationContext) {
	localctx = NewInterfaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, StrictusParserRULE_interfaceDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(StrictusParserInterface)
	}
	{
		p.SetState(209)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(210)
		p.Match(StrictusParserT__3)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(211)
				p.Field()
			}

		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserIdentifier {
		{
			p.SetState(217)
			p.Initializer(false)
		}

	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserPub-34))|(1<<(StrictusParserPubSet-34)))) != 0 {
		{
			p.SetState(220)
			p.FunctionDeclaration(false)
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(226)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IInitializerContext is an interface to support dynamic dispatch.
type IInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsInitializerContext differentiates from other interfaces.
	IsInitializerContext()
}

type InitializerContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	b                     IFunctionBlockContext
}

func NewEmptyInitializerContext() *InitializerContext {
	var p = new(InitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_initializer
	return p
}

func (*InitializerContext) IsInitializerContext() {}

func NewInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *InitializerContext {
	var p = new(InitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_initializer

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *InitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *InitializerContext) GetB() IFunctionBlockContext { return s.b }

func (s *InitializerContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *InitializerContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *InitializerContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *InitializerContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *InitializerContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *InitializerContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *InitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInitializer(s)
	}
}

func (s *InitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInitializer(s)
	}
}

func (s *InitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Initializer(functionBlockRequired bool) (localctx IInitializerContext) {
	localctx = NewInitializerContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 18, StrictusParserRULE_initializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(229)
		p.ParameterList()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(230)

			var _x = p.FunctionBlock()

			localctx.(*InitializerContext).b = _x
		}

	}
	p.SetState(233)

	if !(!localctx.(*InitializerContext).functionBlockRequired || localctx.(*InitializerContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() IFullTypeContext

	// GetB returns the b rule contexts.
	GetB() IFunctionBlockContext

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(IFullTypeContext)

	// SetB sets the b rule contexts.
	SetB(IFunctionBlockContext)

	// GetFunctionBlockRequired returns the functionBlockRequired attribute.
	GetFunctionBlockRequired() bool

	// SetFunctionBlockRequired sets the functionBlockRequired attribute.
	SetFunctionBlockRequired(bool)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	functionBlockRequired bool
	returnType            IFullTypeContext
	b                     IFunctionBlockContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, functionBlockRequired bool) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionDeclaration

	p.functionBlockRequired = functionBlockRequired

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetReturnType() IFullTypeContext { return s.returnType }

func (s *FunctionDeclarationContext) GetB() IFunctionBlockContext { return s.b }

func (s *FunctionDeclarationContext) SetReturnType(v IFullTypeContext) { s.returnType = v }

func (s *FunctionDeclarationContext) SetB(v IFunctionBlockContext) { s.b = v }

func (s *FunctionDeclarationContext) GetFunctionBlockRequired() bool { return s.functionBlockRequired }

func (s *FunctionDeclarationContext) SetFunctionBlockRequired(v bool) { s.functionBlockRequired = v }

func (s *FunctionDeclarationContext) Access() IAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessContext)
}

func (s *FunctionDeclarationContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *FunctionDeclarationContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionDeclaration(functionBlockRequired bool) (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState(), functionBlockRequired)
	p.EnterRule(localctx, 20, StrictusParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Access()
	}
	{
		p.SetState(236)
		p.Match(StrictusParserFun)
	}
	{
		p.SetState(237)
		p.Match(StrictusParserIdentifier)
	}
	{
		p.SetState(238)
		p.ParameterList()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(239)
			p.Match(StrictusParserT__5)
		}
		{
			p.SetState(240)

			var _x = p.FullType()

			localctx.(*FunctionDeclarationContext).returnType = _x
		}

	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(243)

			var _x = p.FunctionBlock()

			localctx.(*FunctionDeclarationContext).b = _x
		}

	}
	p.SetState(246)

	if !(!localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil) {
		panic(antlr.NewFailedPredicateException(p, " !$functionBlockRequired || $ctx.b != nil ", ""))
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *ParameterListContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, StrictusParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserIdentifier {
		{
			p.SetState(249)
			p.Parameter()
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(250)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(251)
				p.Parameter()
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(259)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgumentLabel returns the argumentLabel token.
	GetArgumentLabel() antlr.Token

	// GetParameterName returns the parameterName token.
	GetParameterName() antlr.Token

	// SetArgumentLabel sets the argumentLabel token.
	SetArgumentLabel(antlr.Token)

	// SetParameterName sets the parameterName token.
	SetParameterName(antlr.Token)

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	argumentLabel antlr.Token
	parameterName antlr.Token
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) GetArgumentLabel() antlr.Token { return s.argumentLabel }

func (s *ParameterContext) GetParameterName() antlr.Token { return s.parameterName }

func (s *ParameterContext) SetArgumentLabel(v antlr.Token) { s.argumentLabel = v }

func (s *ParameterContext) SetParameterName(v antlr.Token) { s.parameterName = v }

func (s *ParameterContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *ParameterContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserIdentifier)
}

func (s *ParameterContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, i)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, StrictusParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(261)

			var _m = p.Match(StrictusParserIdentifier)

			localctx.(*ParameterContext).argumentLabel = _m
		}

	}
	{
		p.SetState(264)

		var _m = p.Match(StrictusParserIdentifier)

		localctx.(*ParameterContext).parameterName = _m
	}
	{
		p.SetState(265)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(266)
		p.FullType()
	}

	return localctx
}

// IFullTypeContext is an interface to support dynamic dispatch.
type IFullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_Optional returns the _Optional token.
	Get_Optional() antlr.Token

	// Get_NilCoalescing returns the _NilCoalescing token.
	Get_NilCoalescing() antlr.Token

	// Get_tset322 returns the _tset322 token.
	Get_tset322() antlr.Token

	// Set_Optional sets the _Optional token.
	Set_Optional(antlr.Token)

	// Set_NilCoalescing sets the _NilCoalescing token.
	Set_NilCoalescing(antlr.Token)

	// Set_tset322 sets the _tset322 token.
	Set_tset322(antlr.Token)

	// GetOptionals returns the optionals token list.
	GetOptionals() []antlr.Token

	// SetOptionals sets the optionals token list.
	SetOptionals([]antlr.Token)

	// IsFullTypeContext differentiates from other interfaces.
	IsFullTypeContext()
}

type FullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_Optional      antlr.Token
	optionals      []antlr.Token
	_NilCoalescing antlr.Token
	_tset322       antlr.Token
}

func NewEmptyFullTypeContext() *FullTypeContext {
	var p = new(FullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_fullType
	return p
}

func (*FullTypeContext) IsFullTypeContext() {}

func NewFullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullTypeContext {
	var p = new(FullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_fullType

	return p
}

func (s *FullTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FullTypeContext) Get_Optional() antlr.Token { return s._Optional }

func (s *FullTypeContext) Get_NilCoalescing() antlr.Token { return s._NilCoalescing }

func (s *FullTypeContext) Get_tset322() antlr.Token { return s._tset322 }

func (s *FullTypeContext) Set_Optional(v antlr.Token) { s._Optional = v }

func (s *FullTypeContext) Set_NilCoalescing(v antlr.Token) { s._NilCoalescing = v }

func (s *FullTypeContext) Set_tset322(v antlr.Token) { s._tset322 = v }

func (s *FullTypeContext) GetOptionals() []antlr.Token { return s.optionals }

func (s *FullTypeContext) SetOptionals(v []antlr.Token) { s.optionals = v }

func (s *FullTypeContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *FullTypeContext) AllTypeDimension() []ITypeDimensionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDimensionContext)(nil)).Elem())
	var tst = make([]ITypeDimensionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDimensionContext)
		}
	}

	return tst
}

func (s *FullTypeContext) TypeDimension(i int) ITypeDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDimensionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDimensionContext)
}

func (s *FullTypeContext) AllOptional() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserOptional)
}

func (s *FullTypeContext) Optional(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserOptional, i)
}

func (s *FullTypeContext) AllNilCoalescing() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserNilCoalescing)
}

func (s *FullTypeContext) NilCoalescing(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserNilCoalescing, i)
}

func (s *FullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFullType(s)
	}
}

func (s *FullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFullType(s)
	}
}

func (s *FullTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFullType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FullType() (localctx IFullTypeContext) {
	localctx = NewFullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, StrictusParserRULE_fullType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.BaseType()
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(269)
				p.TypeDimension()
			}

		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(275)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*FullTypeContext)._tset322 = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == StrictusParserOptional || _la == StrictusParserNilCoalescing) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*FullTypeContext)._tset322 = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			localctx.(*FullTypeContext).optionals = append(localctx.(*FullTypeContext).optionals, localctx.(*FullTypeContext)._tset322)

		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeDimensionContext is an interface to support dynamic dispatch.
type ITypeDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDimensionContext differentiates from other interfaces.
	IsTypeDimensionContext()
}

type TypeDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDimensionContext() *TypeDimensionContext {
	var p = new(TypeDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_typeDimension
	return p
}

func (*TypeDimensionContext) IsTypeDimensionContext() {}

func NewTypeDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDimensionContext {
	var p = new(TypeDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_typeDimension

	return p
}

func (s *TypeDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDimensionContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *TypeDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterTypeDimension(s)
	}
}

func (s *TypeDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitTypeDimension(s)
	}
}

func (s *TypeDimensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitTypeDimension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) TypeDimension() (localctx ITypeDimensionContext) {
	localctx = NewTypeDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, StrictusParserRULE_typeDimension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(StrictusParserT__6)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserDecimalLiteral {
		{
			p.SetState(282)
			p.Match(StrictusParserDecimalLiteral)
		}

	}
	{
		p.SetState(285)
		p.Match(StrictusParserT__7)
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *BaseTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, StrictusParserRULE_baseType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(StrictusParserIdentifier)
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.FunctionType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_fullType returns the _fullType rule contexts.
	Get_fullType() IFullTypeContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() IFullTypeContext

	// Set_fullType sets the _fullType rule contexts.
	Set_fullType(IFullTypeContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(IFullTypeContext)

	// GetParameterTypes returns the parameterTypes rule context list.
	GetParameterTypes() []IFullTypeContext

	// SetParameterTypes sets the parameterTypes rule context list.
	SetParameterTypes([]IFullTypeContext)

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_fullType      IFullTypeContext
	parameterTypes []IFullTypeContext
	returnType     IFullTypeContext
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) Get_fullType() IFullTypeContext { return s._fullType }

func (s *FunctionTypeContext) GetReturnType() IFullTypeContext { return s.returnType }

func (s *FunctionTypeContext) Set_fullType(v IFullTypeContext) { s._fullType = v }

func (s *FunctionTypeContext) SetReturnType(v IFullTypeContext) { s.returnType = v }

func (s *FunctionTypeContext) GetParameterTypes() []IFullTypeContext { return s.parameterTypes }

func (s *FunctionTypeContext) SetParameterTypes(v []IFullTypeContext) { s.parameterTypes = v }

func (s *FunctionTypeContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserOpenParen)
}

func (s *FunctionTypeContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, i)
}

func (s *FunctionTypeContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(StrictusParserCloseParen)
}

func (s *FunctionTypeContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, i)
}

func (s *FunctionTypeContext) AllFullType() []IFullTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFullTypeContext)(nil)).Elem())
	var tst = make([]IFullTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFullTypeContext)
		}
	}

	return tst
}

func (s *FunctionTypeContext) FullType(i int) IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, StrictusParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(StrictusParserOpenParen)
	}
	{
		p.SetState(292)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserOpenParen || _la == StrictusParserIdentifier {
		{
			p.SetState(293)

			var _x = p.FullType()

			localctx.(*FunctionTypeContext)._fullType = _x
		}
		localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._fullType)
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(294)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(295)

				var _x = p.FullType()

				localctx.(*FunctionTypeContext)._fullType = _x
			}
			localctx.(*FunctionTypeContext).parameterTypes = append(localctx.(*FunctionTypeContext).parameterTypes, localctx.(*FunctionTypeContext)._fullType)

			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(303)
		p.Match(StrictusParserCloseParen)
	}
	{
		p.SetState(304)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(305)

		var _x = p.FullType()

		localctx.(*FunctionTypeContext).returnType = _x
	}
	{
		p.SetState(306)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, StrictusParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(StrictusParserT__3)
	}
	{
		p.SetState(309)
		p.Statements()
	}
	{
		p.SetState(310)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IFunctionBlockContext is an interface to support dynamic dispatch.
type IFunctionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBlockContext differentiates from other interfaces.
	IsFunctionBlockContext()
}

type FunctionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBlockContext() *FunctionBlockContext {
	var p = new(FunctionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_functionBlock
	return p
}

func (*FunctionBlockContext) IsFunctionBlockContext() {}

func NewFunctionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBlockContext {
	var p = new(FunctionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_functionBlock

	return p
}

func (s *FunctionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBlockContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *FunctionBlockContext) PreConditions() IPreConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreConditionsContext)
}

func (s *FunctionBlockContext) PostConditions() IPostConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostConditionsContext)
}

func (s *FunctionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionBlock(s)
	}
}

func (s *FunctionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FunctionBlock() (localctx IFunctionBlockContext) {
	localctx = NewFunctionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, StrictusParserRULE_functionBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(StrictusParserT__3)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserPre {
		{
			p.SetState(313)
			p.PreConditions()
		}

	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserPost {
		{
			p.SetState(316)
			p.PostConditions()
		}

	}
	{
		p.SetState(319)
		p.Statements()
	}
	{
		p.SetState(320)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IPreConditionsContext is an interface to support dynamic dispatch.
type IPreConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreConditionsContext differentiates from other interfaces.
	IsPreConditionsContext()
}

type PreConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreConditionsContext() *PreConditionsContext {
	var p = new(PreConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_preConditions
	return p
}

func (*PreConditionsContext) IsPreConditionsContext() {}

func NewPreConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreConditionsContext {
	var p = new(PreConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_preConditions

	return p
}

func (s *PreConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PreConditionsContext) Pre() antlr.TerminalNode {
	return s.GetToken(StrictusParserPre, 0)
}

func (s *PreConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PreConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPreConditions(s)
	}
}

func (s *PreConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPreConditions(s)
	}
}

func (s *PreConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPreConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PreConditions() (localctx IPreConditionsContext) {
	localctx = NewPreConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, StrictusParserRULE_preConditions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(StrictusParserPre)
	}
	{
		p.SetState(323)
		p.Match(StrictusParserT__3)
	}
	{
		p.SetState(324)
		p.Conditions()
	}
	{
		p.SetState(325)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IPostConditionsContext is an interface to support dynamic dispatch.
type IPostConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostConditionsContext differentiates from other interfaces.
	IsPostConditionsContext()
}

type PostConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostConditionsContext() *PostConditionsContext {
	var p = new(PostConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_postConditions
	return p
}

func (*PostConditionsContext) IsPostConditionsContext() {}

func NewPostConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostConditionsContext {
	var p = new(PostConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_postConditions

	return p
}

func (s *PostConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PostConditionsContext) Post() antlr.TerminalNode {
	return s.GetToken(StrictusParserPost, 0)
}

func (s *PostConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PostConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPostConditions(s)
	}
}

func (s *PostConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPostConditions(s)
	}
}

func (s *PostConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPostConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PostConditions() (localctx IPostConditionsContext) {
	localctx = NewPostConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, StrictusParserRULE_postConditions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(StrictusParserPost)
	}
	{
		p.SetState(328)
		p.Match(StrictusParserT__3)
	}
	{
		p.SetState(329)
		p.Conditions()
	}
	{
		p.SetState(330)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (s *ConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, StrictusParserRULE_conditions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserT__6)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserTrue-34))|(1<<(StrictusParserFalse-34))|(1<<(StrictusParserNil-34))|(1<<(StrictusParserIdentifier-34))|(1<<(StrictusParserDecimalLiteral-34))|(1<<(StrictusParserBinaryLiteral-34))|(1<<(StrictusParserOctalLiteral-34))|(1<<(StrictusParserHexadecimalLiteral-34))|(1<<(StrictusParserInvalidNumberLiteral-34))|(1<<(StrictusParserStringLiteral-34)))) != 0) {
		{
			p.SetState(332)
			p.Condition()
		}
		{
			p.SetState(333)
			p.Eos()
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTest returns the test rule contexts.
	GetTest() IExpressionContext

	// GetMessage returns the message rule contexts.
	GetMessage() IExpressionContext

	// SetTest sets the test rule contexts.
	SetTest(IExpressionContext)

	// SetMessage sets the message rule contexts.
	SetMessage(IExpressionContext)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	test    IExpressionContext
	message IExpressionContext
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetTest() IExpressionContext { return s.test }

func (s *ConditionContext) GetMessage() IExpressionContext { return s.message }

func (s *ConditionContext) SetTest(v IExpressionContext) { s.test = v }

func (s *ConditionContext) SetMessage(v IExpressionContext) { s.message = v }

func (s *ConditionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, StrictusParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)

		var _x = p.Expression()

		localctx.(*ConditionContext).test = _x
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(341)
			p.Match(StrictusParserT__5)
		}
		{
			p.SetState(342)

			var _x = p.Expression()

			localctx.(*ConditionContext).message = _x
		}

	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementsContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, StrictusParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__0)|(1<<StrictusParserT__3)|(1<<StrictusParserT__6)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(StrictusParserStruct-32))|(1<<(StrictusParserInterface-32))|(1<<(StrictusParserFun-32))|(1<<(StrictusParserPub-32))|(1<<(StrictusParserPubSet-32))|(1<<(StrictusParserReturn-32))|(1<<(StrictusParserBreak-32))|(1<<(StrictusParserContinue-32))|(1<<(StrictusParserLet-32))|(1<<(StrictusParserVar-32))|(1<<(StrictusParserIf-32))|(1<<(StrictusParserWhile-32))|(1<<(StrictusParserTrue-32))|(1<<(StrictusParserFalse-32))|(1<<(StrictusParserNil-32))|(1<<(StrictusParserIdentifier-32))|(1<<(StrictusParserDecimalLiteral-32))|(1<<(StrictusParserBinaryLiteral-32))|(1<<(StrictusParserOctalLiteral-32))|(1<<(StrictusParserHexadecimalLiteral-32))|(1<<(StrictusParserInvalidNumberLiteral-32))|(1<<(StrictusParserStringLiteral-32)))) != 0) {
		{
			p.SetState(345)
			p.Statement()
		}
		{
			p.SetState(346)
			p.Eos()
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, StrictusParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.ReturnStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.BreakStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(355)
			p.ContinueStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(356)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(357)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(358)
			p.Declaration()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(359)
			p.Assignment()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(360)
			p.Expression()
		}

	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(StrictusParserReturn, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, StrictusParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(StrictusParserReturn)
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(364)
			p.Expression()
		}

	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(StrictusParserBreak, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, StrictusParserRULE_breakStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(StrictusParserBreak)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(StrictusParserContinue, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, StrictusParserRULE_continueStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(StrictusParserContinue)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTestExpression returns the testExpression rule contexts.
	GetTestExpression() IExpressionContext

	// GetTestDeclaration returns the testDeclaration rule contexts.
	GetTestDeclaration() IVariableDeclarationContext

	// GetThen returns the then rule contexts.
	GetThen() IBlockContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IBlockContext

	// SetTestExpression sets the testExpression rule contexts.
	SetTestExpression(IExpressionContext)

	// SetTestDeclaration sets the testDeclaration rule contexts.
	SetTestDeclaration(IVariableDeclarationContext)

	// SetThen sets the then rule contexts.
	SetThen(IBlockContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IBlockContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	testExpression  IExpressionContext
	testDeclaration IVariableDeclarationContext
	then            IBlockContext
	alt             IBlockContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetTestExpression() IExpressionContext { return s.testExpression }

func (s *IfStatementContext) GetTestDeclaration() IVariableDeclarationContext {
	return s.testDeclaration
}

func (s *IfStatementContext) GetThen() IBlockContext { return s.then }

func (s *IfStatementContext) GetAlt() IBlockContext { return s.alt }

func (s *IfStatementContext) SetTestExpression(v IExpressionContext) { s.testExpression = v }

func (s *IfStatementContext) SetTestDeclaration(v IVariableDeclarationContext) { s.testDeclaration = v }

func (s *IfStatementContext) SetThen(v IBlockContext) { s.then = v }

func (s *IfStatementContext) SetAlt(v IBlockContext) { s.alt = v }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(StrictusParserIf, 0)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(StrictusParserElse, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, StrictusParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(StrictusParserIf)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__3, StrictusParserT__6, StrictusParserMinus, StrictusParserNegate, StrictusParserOpenParen, StrictusParserFun, StrictusParserTrue, StrictusParserFalse, StrictusParserNil, StrictusParserIdentifier, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral, StrictusParserStringLiteral:
		{
			p.SetState(372)

			var _x = p.Expression()

			localctx.(*IfStatementContext).testExpression = _x
		}

	case StrictusParserLet, StrictusParserVar:
		{
			p.SetState(373)

			var _x = p.VariableDeclaration()

			localctx.(*IfStatementContext).testDeclaration = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(376)

		var _x = p.Block()

		localctx.(*IfStatementContext).then = _x
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(377)
			p.Match(StrictusParserElse)
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case StrictusParserIf:
			{
				p.SetState(378)
				p.IfStatement()
			}

		case StrictusParserT__3:
			{
				p.SetState(379)

				var _x = p.Block()

				localctx.(*IfStatementContext).alt = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(StrictusParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, StrictusParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(StrictusParserWhile)
	}
	{
		p.SetState(385)
		p.Expression()
	}
	{
		p.SetState(386)
		p.Block()
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableKind() IVariableKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableKindContext)
}

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, StrictusParserRULE_variableDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.VariableKind()
	}
	{
		p.SetState(389)
		p.Match(StrictusParserIdentifier)
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == StrictusParserT__5 {
		{
			p.SetState(390)
			p.Match(StrictusParserT__5)
		}
		{
			p.SetState(391)
			p.FullType()
		}

	}
	{
		p.SetState(394)
		p.Match(StrictusParserT__8)
	}
	{
		p.SetState(395)
		p.Expression()
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) AllExpressionAccess() []IExpressionAccessContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem())
	var tst = make([]IExpressionAccessContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAccessContext)
		}
	}

	return tst
}

func (s *AssignmentContext) ExpressionAccess(i int) IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, StrictusParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(StrictusParserIdentifier)
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == StrictusParserT__6 || _la == StrictusParserT__11 {
		{
			p.SetState(398)
			p.ExpressionAccess()
		}

		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(404)
		p.Match(StrictusParserT__8)
	}
	{
		p.SetState(405)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, StrictusParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.ConditionalExpression()
	}

	return localctx
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetThen returns the then rule contexts.
	GetThen() IExpressionContext

	// GetAlt returns the alt rule contexts.
	GetAlt() IExpressionContext

	// SetThen sets the then rule contexts.
	SetThen(IExpressionContext)

	// SetAlt sets the alt rule contexts.
	SetAlt(IExpressionContext)

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	then   IExpressionContext
	alt    IExpressionContext
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_conditionalExpression
	return p
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) GetThen() IExpressionContext { return s.then }

func (s *ConditionalExpressionContext) GetAlt() IExpressionContext { return s.alt }

func (s *ConditionalExpressionContext) SetThen(v IExpressionContext) { s.then = v }

func (s *ConditionalExpressionContext) SetAlt(v IExpressionContext) { s.alt = v }

func (s *ConditionalExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ConditionalExpressionContext) Optional() antlr.TerminalNode {
	return s.GetToken(StrictusParserOptional, 0)
}

func (s *ConditionalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitConditionalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, StrictusParserRULE_conditionalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.orExpression(0)
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(410)
			p.Match(StrictusParserOptional)
		}
		{
			p.SetState(411)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).then = _x
		}
		{
			p.SetState(412)
			p.Match(StrictusParserT__5)
		}
		{
			p.SetState(413)

			var _x = p.Expression()

			localctx.(*ConditionalExpressionContext).alt = _x
		}

	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) OrExpression() (localctx IOrExpressionContext) {
	return p.orExpression(0)
}

func (p *StrictusParser) orExpression(_p int) (localctx IOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, StrictusParserRULE_orExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_orExpression)
			p.SetState(420)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(421)
				p.Match(StrictusParserT__9)
			}
			{
				p.SetState(422)
				p.andExpression(0)
			}

		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *StrictusParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, StrictusParserRULE_andExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_andExpression)
			p.SetState(431)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(432)
				p.Match(StrictusParserT__10)
			}
			{
				p.SetState(433)
				p.equalityExpression(0)
			}

		}
		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) EqualityOp() IEqualityOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityOpContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *StrictusParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 72
	p.EnterRecursionRule(localctx, 72, StrictusParserRULE_equalityExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_equalityExpression)
			p.SetState(442)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(443)
				p.EqualityOp()
			}
			{
				p.SetState(444)
				p.relationalExpression(0)
			}

		}
		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) RelationalOp() IRelationalOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalOpContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *StrictusParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, StrictusParserRULE_relationalExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.NilCoalescingExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_relationalExpression)
			p.SetState(454)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(455)
				p.RelationalOp()
			}
			{
				p.SetState(456)
				p.NilCoalescingExpression()
			}

		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// INilCoalescingExpressionContext is an interface to support dynamic dispatch.
type INilCoalescingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilCoalescingExpressionContext differentiates from other interfaces.
	IsNilCoalescingExpressionContext()
}

type NilCoalescingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilCoalescingExpressionContext() *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nilCoalescingExpression
	return p
}

func (*NilCoalescingExpressionContext) IsNilCoalescingExpressionContext() {}

func NewNilCoalescingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilCoalescingExpressionContext {
	var p = new(NilCoalescingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nilCoalescingExpression

	return p
}

func (s *NilCoalescingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NilCoalescingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *NilCoalescingExpressionContext) NilCoalescing() antlr.TerminalNode {
	return s.GetToken(StrictusParserNilCoalescing, 0)
}

func (s *NilCoalescingExpressionContext) NilCoalescingExpression() INilCoalescingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilCoalescingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilCoalescingExpressionContext)
}

func (s *NilCoalescingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilCoalescingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilCoalescingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNilCoalescingExpression(s)
	}
}

func (s *NilCoalescingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNilCoalescingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NilCoalescingExpression() (localctx INilCoalescingExpressionContext) {
	localctx = NewNilCoalescingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, StrictusParserRULE_nilCoalescingExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		p.failableDowncastingExpression(0)
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(464)
			p.Match(StrictusParserNilCoalescing)
		}
		{
			p.SetState(465)
			p.NilCoalescingExpression()
		}

	}

	return localctx
}

// IFailableDowncastingExpressionContext is an interface to support dynamic dispatch.
type IFailableDowncastingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFailableDowncastingExpressionContext differentiates from other interfaces.
	IsFailableDowncastingExpressionContext()
}

type FailableDowncastingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFailableDowncastingExpressionContext() *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_failableDowncastingExpression
	return p
}

func (*FailableDowncastingExpressionContext) IsFailableDowncastingExpressionContext() {}

func NewFailableDowncastingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FailableDowncastingExpressionContext {
	var p = new(FailableDowncastingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_failableDowncastingExpression

	return p
}

func (s *FailableDowncastingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FailableDowncastingExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncastingExpression() IFailableDowncastingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFailableDowncastingExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFailableDowncastingExpressionContext)
}

func (s *FailableDowncastingExpressionContext) FailableDowncasting() antlr.TerminalNode {
	return s.GetToken(StrictusParserFailableDowncasting, 0)
}

func (s *FailableDowncastingExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *FailableDowncastingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailableDowncastingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FailableDowncastingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFailableDowncastingExpression(s)
	}
}

func (s *FailableDowncastingExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFailableDowncastingExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) FailableDowncastingExpression() (localctx IFailableDowncastingExpressionContext) {
	return p.failableDowncastingExpression(0)
}

func (p *StrictusParser) failableDowncastingExpression(_p int) (localctx IFailableDowncastingExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFailableDowncastingExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFailableDowncastingExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, StrictusParserRULE_failableDowncastingExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.additiveExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFailableDowncastingExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_failableDowncastingExpression)
			p.SetState(471)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(472)
				p.Match(StrictusParserFailableDowncasting)
			}
			{
				p.SetState(473)
				p.FullType()
			}

		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *AdditiveExpressionContext) AdditiveOp() IAdditiveOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditiveOpContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	return p.additiveExpression(0)
}

func (p *StrictusParser) additiveExpression(_p int) (localctx IAdditiveExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditiveExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, StrictusParserRULE_additiveExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.multiplicativeExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAdditiveExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_additiveExpression)
			p.SetState(482)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(483)
				p.AdditiveOp()
			}
			{
				p.SetState(484)
				p.multiplicativeExpression(0)
			}

		}
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeExpression() IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *MultiplicativeExpressionContext) MultiplicativeOp() IMultiplicativeOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOpContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	return p.multiplicativeExpression(0)
}

func (p *StrictusParser) multiplicativeExpression(_p int) (localctx IMultiplicativeExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicativeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, StrictusParserRULE_multiplicativeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.UnaryExpression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMultiplicativeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, StrictusParserRULE_multiplicativeExpression)
			p.SetState(494)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(495)
				p.MultiplicativeOp()
			}
			{
				p.SetState(496)
				p.UnaryExpression()
			}

		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *UnaryExpressionContext) AllUnaryOp() []IUnaryOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem())
	var tst = make([]IUnaryOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryOpContext)
		}
	}

	return tst
}

func (s *UnaryExpressionContext) UnaryOp(i int) IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, StrictusParserRULE_unaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(511)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__3, StrictusParserT__6, StrictusParserOpenParen, StrictusParserFun, StrictusParserTrue, StrictusParserFalse, StrictusParserNil, StrictusParserIdentifier, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral, StrictusParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(503)
			p.PrimaryExpression()
		}

	case StrictusParserMinus, StrictusParserNegate:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(504)
					p.UnaryOp()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(507)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
		}
		{
			p.SetState(509)
			p.UnaryExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) PrimaryExpressionStart() IPrimaryExpressionStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionStartContext)
}

func (s *PrimaryExpressionContext) AllPrimaryExpressionSuffix() []IPrimaryExpressionSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem())
	var tst = make([]IPrimaryExpressionSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryExpressionSuffixContext)
		}
	}

	return tst
}

func (s *PrimaryExpressionContext) PrimaryExpressionSuffix(i int) IPrimaryExpressionSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionSuffixContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, StrictusParserRULE_primaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.PrimaryExpressionStart()
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(514)
				p.PrimaryExpressionSuffix()
			}

		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionSuffixContext is an interface to support dynamic dispatch.
type IPrimaryExpressionSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionSuffixContext differentiates from other interfaces.
	IsPrimaryExpressionSuffixContext()
}

type PrimaryExpressionSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionSuffixContext() *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix
	return p
}

func (*PrimaryExpressionSuffixContext) IsPrimaryExpressionSuffixContext() {}

func NewPrimaryExpressionSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionSuffixContext {
	var p = new(PrimaryExpressionSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionSuffix

	return p
}

func (s *PrimaryExpressionSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionSuffixContext) ExpressionAccess() IExpressionAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAccessContext)
}

func (s *PrimaryExpressionSuffixContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *PrimaryExpressionSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitPrimaryExpressionSuffix(s)
	}
}

func (s *PrimaryExpressionSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitPrimaryExpressionSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionSuffix() (localctx IPrimaryExpressionSuffixContext) {
	localctx = NewPrimaryExpressionSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, StrictusParserRULE_primaryExpressionSuffix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(522)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__6, StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(520)
			p.ExpressionAccess()
		}

	case StrictusParserOpenParen:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(521)
			p.Invocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqualityOpContext is an interface to support dynamic dispatch.
type IEqualityOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityOpContext differentiates from other interfaces.
	IsEqualityOpContext()
}

type EqualityOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOpContext() *EqualityOpContext {
	var p = new(EqualityOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_equalityOp
	return p
}

func (*EqualityOpContext) IsEqualityOpContext() {}

func NewEqualityOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOpContext {
	var p = new(EqualityOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_equalityOp

	return p
}

func (s *EqualityOpContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityOpContext) Equal() antlr.TerminalNode {
	return s.GetToken(StrictusParserEqual, 0)
}

func (s *EqualityOpContext) Unequal() antlr.TerminalNode {
	return s.GetToken(StrictusParserUnequal, 0)
}

func (s *EqualityOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEqualityOp(s)
	}
}

func (s *EqualityOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEqualityOp(s)
	}
}

func (s *EqualityOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEqualityOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) EqualityOp() (localctx IEqualityOpContext) {
	localctx = NewEqualityOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, StrictusParserRULE_equalityOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserEqual || _la == StrictusParserUnequal) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelationalOpContext is an interface to support dynamic dispatch.
type IRelationalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalOpContext differentiates from other interfaces.
	IsRelationalOpContext()
}

type RelationalOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOpContext() *RelationalOpContext {
	var p = new(RelationalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_relationalOp
	return p
}

func (*RelationalOpContext) IsRelationalOpContext() {}

func NewRelationalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOpContext {
	var p = new(RelationalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_relationalOp

	return p
}

func (s *RelationalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOpContext) Less() antlr.TerminalNode {
	return s.GetToken(StrictusParserLess, 0)
}

func (s *RelationalOpContext) Greater() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreater, 0)
}

func (s *RelationalOpContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserLessEqual, 0)
}

func (s *RelationalOpContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(StrictusParserGreaterEqual, 0)
}

func (s *RelationalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterRelationalOp(s)
	}
}

func (s *RelationalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitRelationalOp(s)
	}
}

func (s *RelationalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitRelationalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) RelationalOp() (localctx IRelationalOpContext) {
	localctx = NewRelationalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, StrictusParserRULE_relationalOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserLess)|(1<<StrictusParserGreater)|(1<<StrictusParserLessEqual)|(1<<StrictusParserGreaterEqual))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAdditiveOpContext is an interface to support dynamic dispatch.
type IAdditiveOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveOpContext differentiates from other interfaces.
	IsAdditiveOpContext()
}

type AdditiveOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOpContext() *AdditiveOpContext {
	var p = new(AdditiveOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_additiveOp
	return p
}

func (*AdditiveOpContext) IsAdditiveOpContext() {}

func NewAdditiveOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOpContext {
	var p = new(AdditiveOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_additiveOp

	return p
}

func (s *AdditiveOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOpContext) Plus() antlr.TerminalNode {
	return s.GetToken(StrictusParserPlus, 0)
}

func (s *AdditiveOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *AdditiveOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitAdditiveOp(s)
	}
}

func (s *AdditiveOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitAdditiveOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) AdditiveOp() (localctx IAdditiveOpContext) {
	localctx = NewAdditiveOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, StrictusParserRULE_additiveOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserPlus || _la == StrictusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultiplicativeOpContext is an interface to support dynamic dispatch.
type IMultiplicativeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeOpContext differentiates from other interfaces.
	IsMultiplicativeOpContext()
}

type MultiplicativeOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOpContext() *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_multiplicativeOp
	return p
}

func (*MultiplicativeOpContext) IsMultiplicativeOpContext() {}

func NewMultiplicativeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOpContext {
	var p = new(MultiplicativeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_multiplicativeOp

	return p
}

func (s *MultiplicativeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOpContext) Mul() antlr.TerminalNode {
	return s.GetToken(StrictusParserMul, 0)
}

func (s *MultiplicativeOpContext) Div() antlr.TerminalNode {
	return s.GetToken(StrictusParserDiv, 0)
}

func (s *MultiplicativeOpContext) Mod() antlr.TerminalNode {
	return s.GetToken(StrictusParserMod, 0)
}

func (s *MultiplicativeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMultiplicativeOp(s)
	}
}

func (s *MultiplicativeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMultiplicativeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MultiplicativeOp() (localctx IMultiplicativeOpContext) {
	localctx = NewMultiplicativeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, StrictusParserRULE_multiplicativeOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserMul)|(1<<StrictusParserDiv)|(1<<StrictusParserMod))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(StrictusParserMinus, 0)
}

func (s *UnaryOpContext) Negate() antlr.TerminalNode {
	return s.GetToken(StrictusParserNegate, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, StrictusParserRULE_unaryOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserMinus || _la == StrictusParserNegate) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrimaryExpressionStartContext is an interface to support dynamic dispatch.
type IPrimaryExpressionStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionStartContext differentiates from other interfaces.
	IsPrimaryExpressionStartContext()
}

type PrimaryExpressionStartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionStartContext() *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart
	return p
}

func (*PrimaryExpressionStartContext) IsPrimaryExpressionStartContext() {}

func NewPrimaryExpressionStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionStartContext {
	var p = new(PrimaryExpressionStartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_primaryExpressionStart

	return p
}

func (s *PrimaryExpressionStartContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionStartContext) CopyFrom(ctx *PrimaryExpressionStartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionExpressionContext struct {
	*PrimaryExpressionStartContext
	returnType IFullTypeContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *FunctionExpressionContext) GetReturnType() IFullTypeContext { return s.returnType }

func (s *FunctionExpressionContext) SetReturnType(v IFullTypeContext) { s.returnType = v }

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Fun() antlr.TerminalNode {
	return s.GetToken(StrictusParserFun, 0)
}

func (s *FunctionExpressionContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionExpressionContext) FunctionBlock() IFunctionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBlockContext)
}

func (s *FunctionExpressionContext) FullType() IFullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullTypeContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestedExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewNestedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExpressionContext {
	var p = new(NestedExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *NestedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *NestedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *NestedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNestedExpression(s)
	}
}

func (s *NestedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNestedExpression(s)
	}
}

func (s *NestedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNestedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*PrimaryExpressionStartContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.PrimaryExpressionStartContext = NewEmptyPrimaryExpressionStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionStartContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) PrimaryExpressionStart() (localctx IPrimaryExpressionStartContext) {
	localctx = NewPrimaryExpressionStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, StrictusParserRULE_primaryExpressionStart)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(548)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Match(StrictusParserIdentifier)
		}

	case StrictusParserT__3, StrictusParserT__6, StrictusParserTrue, StrictusParserFalse, StrictusParserNil, StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral, StrictusParserStringLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)
			p.Literal()
		}

	case StrictusParserFun:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(536)
			p.Match(StrictusParserFun)
		}
		{
			p.SetState(537)
			p.ParameterList()
		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == StrictusParserT__5 {
			{
				p.SetState(538)
				p.Match(StrictusParserT__5)
			}
			{
				p.SetState(539)

				var _x = p.FullType()

				localctx.(*FunctionExpressionContext).returnType = _x
			}

		}
		{
			p.SetState(542)
			p.FunctionBlock()
		}

	case StrictusParserOpenParen:
		localctx = NewNestedExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(544)
			p.Match(StrictusParserOpenParen)
		}
		{
			p.SetState(545)
			p.Expression()
		}
		{
			p.SetState(546)
			p.Match(StrictusParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionAccessContext is an interface to support dynamic dispatch.
type IExpressionAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAccessContext differentiates from other interfaces.
	IsExpressionAccessContext()
}

type ExpressionAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAccessContext() *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_expressionAccess
	return p
}

func (*ExpressionAccessContext) IsExpressionAccessContext() {}

func NewExpressionAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAccessContext {
	var p = new(ExpressionAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_expressionAccess

	return p
}

func (s *ExpressionAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAccessContext) MemberAccess() IMemberAccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberAccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberAccessContext)
}

func (s *ExpressionAccessContext) BracketExpression() IBracketExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketExpressionContext)
}

func (s *ExpressionAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitExpressionAccess(s)
	}
}

func (s *ExpressionAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitExpressionAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ExpressionAccess() (localctx IExpressionAccessContext) {
	localctx = NewExpressionAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, StrictusParserRULE_expressionAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(552)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(550)
			p.MemberAccess()
		}

	case StrictusParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(551)
			p.BracketExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberAccessContext is an interface to support dynamic dispatch.
type IMemberAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberAccessContext differentiates from other interfaces.
	IsMemberAccessContext()
}

type MemberAccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAccessContext() *MemberAccessContext {
	var p = new(MemberAccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_memberAccess
	return p
}

func (*MemberAccessContext) IsMemberAccessContext() {}

func NewMemberAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAccessContext {
	var p = new(MemberAccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_memberAccess

	return p
}

func (s *MemberAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAccessContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *MemberAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterMemberAccess(s)
	}
}

func (s *MemberAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitMemberAccess(s)
	}
}

func (s *MemberAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitMemberAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) MemberAccess() (localctx IMemberAccessContext) {
	localctx = NewMemberAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, StrictusParserRULE_memberAccess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(554)
		p.Match(StrictusParserT__11)
	}
	{
		p.SetState(555)
		p.Match(StrictusParserIdentifier)
	}

	return localctx
}

// IBracketExpressionContext is an interface to support dynamic dispatch.
type IBracketExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketExpressionContext differentiates from other interfaces.
	IsBracketExpressionContext()
}

type BracketExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketExpressionContext() *BracketExpressionContext {
	var p = new(BracketExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_bracketExpression
	return p
}

func (*BracketExpressionContext) IsBracketExpressionContext() {}

func NewBracketExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_bracketExpression

	return p
}

func (s *BracketExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

func (s *BracketExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBracketExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BracketExpression() (localctx IBracketExpressionContext) {
	localctx = NewBracketExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, StrictusParserRULE_bracketExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.Match(StrictusParserT__6)
	}
	{
		p.SetState(558)
		p.Expression()
	}
	{
		p.SetState(559)
		p.Match(StrictusParserT__7)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserOpenParen, 0)
}

func (s *InvocationContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(StrictusParserCloseParen, 0)
}

func (s *InvocationContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *InvocationContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (s *InvocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInvocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, StrictusParserRULE_invocation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.Match(StrictusParserOpenParen)
	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserT__6)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserTrue-34))|(1<<(StrictusParserFalse-34))|(1<<(StrictusParserNil-34))|(1<<(StrictusParserIdentifier-34))|(1<<(StrictusParserDecimalLiteral-34))|(1<<(StrictusParserBinaryLiteral-34))|(1<<(StrictusParserOctalLiteral-34))|(1<<(StrictusParserHexadecimalLiteral-34))|(1<<(StrictusParserInvalidNumberLiteral-34))|(1<<(StrictusParserStringLiteral-34)))) != 0) {
		{
			p.SetState(562)
			p.Argument()
		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(563)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(564)
				p.Argument()
			}

			p.SetState(569)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(572)
		p.Match(StrictusParserCloseParen)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(StrictusParserIdentifier, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, StrictusParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(576)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(574)
			p.Match(StrictusParserIdentifier)
		}
		{
			p.SetState(575)
			p.Match(StrictusParserT__5)
		}

	}
	{
		p.SetState(578)
		p.Expression()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) DictionaryLiteral() IDictionaryLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictionaryLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) NilLiteral() INilLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INilLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INilLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, StrictusParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(586)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserDecimalLiteral, StrictusParserBinaryLiteral, StrictusParserOctalLiteral, StrictusParserHexadecimalLiteral, StrictusParserInvalidNumberLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(580)
			p.IntegerLiteral()
		}

	case StrictusParserTrue, StrictusParserFalse:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(581)
			p.BooleanLiteral()
		}

	case StrictusParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(582)
			p.ArrayLiteral()
		}

	case StrictusParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(583)
			p.DictionaryLiteral()
		}

	case StrictusParserStringLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(584)
			p.StringLiteral()
		}

	case StrictusParserNil:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(585)
			p.NilLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) True() antlr.TerminalNode {
	return s.GetToken(StrictusParserTrue, 0)
}

func (s *BooleanLiteralContext) False() antlr.TerminalNode {
	return s.GetToken(StrictusParserFalse, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, StrictusParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(588)
		_la = p.GetTokenStream().LA(1)

		if !(_la == StrictusParserTrue || _la == StrictusParserFalse) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INilLiteralContext is an interface to support dynamic dispatch.
type INilLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilLiteralContext differentiates from other interfaces.
	IsNilLiteralContext()
}

type NilLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilLiteralContext() *NilLiteralContext {
	var p = new(NilLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_nilLiteral
	return p
}

func (*NilLiteralContext) IsNilLiteralContext() {}

func NewNilLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilLiteralContext {
	var p = new(NilLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_nilLiteral

	return p
}

func (s *NilLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NilLiteralContext) Nil() antlr.TerminalNode {
	return s.GetToken(StrictusParserNil, 0)
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) NilLiteral() (localctx INilLiteralContext) {
	localctx = NewNilLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, StrictusParserRULE_nilLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
		p.Match(StrictusParserNil)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserStringLiteral, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, StrictusParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(592)
		p.Match(StrictusParserStringLiteral)
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) CopyFrom(ctx *IntegerLiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryLiteralContext struct {
	*IntegerLiteralContext
}

func NewBinaryLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLiteralContext {
	var p = new(BinaryLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *BinaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLiteralContext) BinaryLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserBinaryLiteral, 0)
}

func (s *BinaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitBinaryLiteral(s)
	}
}

func (s *BinaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitBinaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type OctalLiteralContext struct {
	*IntegerLiteralContext
}

func NewOctalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctalLiteralContext {
	var p = new(OctalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *OctalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserOctalLiteral, 0)
}

func (s *OctalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitOctalLiteral(s)
	}
}

func (s *OctalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitOctalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvalidNumberLiteralContext struct {
	*IntegerLiteralContext
}

func NewInvalidNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvalidNumberLiteralContext {
	var p = new(InvalidNumberLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *InvalidNumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvalidNumberLiteralContext) InvalidNumberLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserInvalidNumberLiteral, 0)
}

func (s *InvalidNumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitInvalidNumberLiteral(s)
	}
}

func (s *InvalidNumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitInvalidNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecimalLiteralContext struct {
	*IntegerLiteralContext
}

func NewDecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserDecimalLiteral, 0)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type HexadecimalLiteralContext struct {
	*IntegerLiteralContext
}

func NewHexadecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexadecimalLiteralContext {
	var p = new(HexadecimalLiteralContext)

	p.IntegerLiteralContext = NewEmptyIntegerLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IntegerLiteralContext))

	return p
}

func (s *HexadecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexadecimalLiteralContext) HexadecimalLiteral() antlr.TerminalNode {
	return s.GetToken(StrictusParserHexadecimalLiteral, 0)
}

func (s *HexadecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitHexadecimalLiteral(s)
	}
}

func (s *HexadecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitHexadecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, StrictusParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(599)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case StrictusParserDecimalLiteral:
		localctx = NewDecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(594)
			p.Match(StrictusParserDecimalLiteral)
		}

	case StrictusParserBinaryLiteral:
		localctx = NewBinaryLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(StrictusParserBinaryLiteral)
		}

	case StrictusParserOctalLiteral:
		localctx = NewOctalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(596)
			p.Match(StrictusParserOctalLiteral)
		}

	case StrictusParserHexadecimalLiteral:
		localctx = NewHexadecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(597)
			p.Match(StrictusParserHexadecimalLiteral)
		}

	case StrictusParserInvalidNumberLiteral:
		localctx = NewInvalidNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(598)
			p.Match(StrictusParserInvalidNumberLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, StrictusParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		p.Match(StrictusParserT__6)
	}
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserT__6)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserTrue-34))|(1<<(StrictusParserFalse-34))|(1<<(StrictusParserNil-34))|(1<<(StrictusParserIdentifier-34))|(1<<(StrictusParserDecimalLiteral-34))|(1<<(StrictusParserBinaryLiteral-34))|(1<<(StrictusParserOctalLiteral-34))|(1<<(StrictusParserHexadecimalLiteral-34))|(1<<(StrictusParserInvalidNumberLiteral-34))|(1<<(StrictusParserStringLiteral-34)))) != 0) {
		{
			p.SetState(602)
			p.Expression()
		}
		p.SetState(607)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(603)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(604)
				p.Expression()
			}

			p.SetState(609)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(612)
		p.Match(StrictusParserT__7)
	}

	return localctx
}

// IDictionaryLiteralContext is an interface to support dynamic dispatch.
type IDictionaryLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictionaryLiteralContext differentiates from other interfaces.
	IsDictionaryLiteralContext()
}

type DictionaryLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictionaryLiteralContext() *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_dictionaryLiteral
	return p
}

func (*DictionaryLiteralContext) IsDictionaryLiteralContext() {}

func NewDictionaryLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryLiteralContext {
	var p = new(DictionaryLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_dictionaryLiteral

	return p
}

func (s *DictionaryLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryLiteralContext) AllDictionaryEntry() []IDictionaryEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem())
	var tst = make([]IDictionaryEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDictionaryEntryContext)
		}
	}

	return tst
}

func (s *DictionaryLiteralContext) DictionaryEntry(i int) IDictionaryEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictionaryEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDictionaryEntryContext)
}

func (s *DictionaryLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDictionaryLiteral(s)
	}
}

func (s *DictionaryLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDictionaryLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DictionaryLiteral() (localctx IDictionaryLiteralContext) {
	localctx = NewDictionaryLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, StrictusParserRULE_dictionaryLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(614)
		p.Match(StrictusParserT__3)
	}
	p.SetState(623)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<StrictusParserT__3)|(1<<StrictusParserT__6)|(1<<StrictusParserMinus)|(1<<StrictusParserNegate)|(1<<StrictusParserOpenParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(StrictusParserFun-34))|(1<<(StrictusParserTrue-34))|(1<<(StrictusParserFalse-34))|(1<<(StrictusParserNil-34))|(1<<(StrictusParserIdentifier-34))|(1<<(StrictusParserDecimalLiteral-34))|(1<<(StrictusParserBinaryLiteral-34))|(1<<(StrictusParserOctalLiteral-34))|(1<<(StrictusParserHexadecimalLiteral-34))|(1<<(StrictusParserInvalidNumberLiteral-34))|(1<<(StrictusParserStringLiteral-34)))) != 0) {
		{
			p.SetState(615)
			p.DictionaryEntry()
		}
		p.SetState(620)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == StrictusParserT__1 {
			{
				p.SetState(616)
				p.Match(StrictusParserT__1)
			}
			{
				p.SetState(617)
				p.DictionaryEntry()
			}

			p.SetState(622)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(625)
		p.Match(StrictusParserT__4)
	}

	return localctx
}

// IDictionaryEntryContext is an interface to support dynamic dispatch.
type IDictionaryEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key rule contexts.
	GetKey() IExpressionContext

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetKey sets the key rule contexts.
	SetKey(IExpressionContext)

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsDictionaryEntryContext differentiates from other interfaces.
	IsDictionaryEntryContext()
}

type DictionaryEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    IExpressionContext
	value  IExpressionContext
}

func NewEmptyDictionaryEntryContext() *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_dictionaryEntry
	return p
}

func (*DictionaryEntryContext) IsDictionaryEntryContext() {}

func NewDictionaryEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictionaryEntryContext {
	var p = new(DictionaryEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_dictionaryEntry

	return p
}

func (s *DictionaryEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictionaryEntryContext) GetKey() IExpressionContext { return s.key }

func (s *DictionaryEntryContext) GetValue() IExpressionContext { return s.value }

func (s *DictionaryEntryContext) SetKey(v IExpressionContext) { s.key = v }

func (s *DictionaryEntryContext) SetValue(v IExpressionContext) { s.value = v }

func (s *DictionaryEntryContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DictionaryEntryContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DictionaryEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictionaryEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictionaryEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitDictionaryEntry(s)
	}
}

func (s *DictionaryEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitDictionaryEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) DictionaryEntry() (localctx IDictionaryEntryContext) {
	localctx = NewDictionaryEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, StrictusParserRULE_dictionaryEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(627)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).key = _x
	}
	{
		p.SetState(628)
		p.Match(StrictusParserT__5)
	}
	{
		p.SetState(629)

		var _x = p.Expression()

		localctx.(*DictionaryEntryContext).value = _x
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrictusParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrictusParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrictusParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrictusListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case StrictusVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *StrictusParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, StrictusParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(631)
			p.Match(StrictusParserT__12)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(632)
			p.Match(StrictusParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(633)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(634)

		if !(p.GetTokenStream().LT(1).GetText() == "}") {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"}\"", ""))
		}

	}

	return localctx
}

func (p *StrictusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *InitializerContext = nil
		if localctx != nil {
			t = localctx.(*InitializerContext)
		}
		return p.Initializer_Sempred(t, predIndex)

	case 10:
		var t *FunctionDeclarationContext = nil
		if localctx != nil {
			t = localctx.(*FunctionDeclarationContext)
		}
		return p.FunctionDeclaration_Sempred(t, predIndex)

	case 34:
		var t *OrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*OrExpressionContext)
		}
		return p.OrExpression_Sempred(t, predIndex)

	case 35:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 36:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 37:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 39:
		var t *FailableDowncastingExpressionContext = nil
		if localctx != nil {
			t = localctx.(*FailableDowncastingExpressionContext)
		}
		return p.FailableDowncastingExpression_Sempred(t, predIndex)

	case 40:
		var t *AdditiveExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveExpressionContext)
		}
		return p.AdditiveExpression_Sempred(t, predIndex)

	case 41:
		var t *MultiplicativeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeExpressionContext)
		}
		return p.MultiplicativeExpression_Sempred(t, predIndex)

	case 64:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *StrictusParser) Initializer_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !localctx.(*InitializerContext).functionBlockRequired || localctx.(*InitializerContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) FunctionDeclaration_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return !localctx.(*FunctionDeclarationContext).functionBlockRequired || localctx.(*FunctionDeclarationContext).b != nil

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) OrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) FailableDowncastingExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) AdditiveExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) MultiplicativeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *StrictusParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.lineTerminatorAhead()

	case 10:
		return p.GetTokenStream().LT(1).GetText() == "}"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
