package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

var initialField = Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
var pieces = []string{"I", "J", "L", "O", "S", "T", "Z"}

var g1 = [300]string{"L", "S", "I", "Z", "Z", "L", "Z", "L", "J", "J", "T", "O", "T", "J", "J", "O", "L", "Z", "I", "Z", "S", "O", "I", "I", "J", "I", "S", "S", "Z", "Z", "S", "T", "I", "I", "O", "Z", "I", "T", "J", "T", "L", "L", "J", "T", "T", "I", "L", "L", "S", "S", "I", "I", "O", "J", "O", "O", "I", "I", "J", "O", "L", "J", "S", "S", "I", "O", "S", "J", "L", "J", "L", "T", "Z", "S", "S", "J", "Z", "J", "Z", "J", "L", "O", "Z", "O", "Z", "I", "Z", "S", "J", "O", "O", "S", "O", "S", "L", "L", "O", "T", "O", "Z", "I", "T", "L", "L", "T", "L", "I", "J", "J", "Z", "O", "O", "Z", "O", "J", "J", "I", "J", "Z", "I", "T", "S", "O", "S", "Z", "L", "Z", "I", "O", "O", "T", "S", "J", "O", "J", "T", "Z", "L", "I", "Z", "Z", "J", "S", "O", "O", "O", "I", "Z", "J", "I", "S", "T", "O", "I", "O", "O", "Z", "S", "I", "Z", "J", "J", "J", "J", "S", "J", "I", "Z", "T", "L", "T", "Z", "S", "T", "O", "Z", "Z", "T", "T", "S", "I", "T", "T", "O", "J", "O", "T", "I", "Z", "O", "J", "S", "T", "T", "S", "O", "S", "L", "Z", "L", "O", "L", "T", "L", "O", "S", "I", "L", "Z", "T", "I", "I", "T", "O", "L", "O", "S", "T", "S", "J", "J", "I", "S", "S", "O", "S", "S", "Z", "L", "I", "Z", "S", "I", "T", "S", "Z", "T", "T", "J", "S", "L", "J", "L", "Z", "I", "S", "T", "T", "O", "T", "I", "I", "S", "I", "O", "I", "O", "T", "J", "I", "T", "T", "Z", "I", "L", "S", "T", "L", "T", "J", "Z", "L", "T", "I", "Z", "I", "Z", "O", "L", "L", "O", "Z", "O", "O", "O", "Z", "T", "I", "J", "S", "I", "I", "S", "S", "I", "O", "Z", "O", "T", "L"}
var g2 = [300]string{"T", "L", "O", "S", "S", "L", "O", "S", "S", "L", "S", "J", "Z", "J", "O", "I", "T", "T", "Z", "L", "S", "O", "I", "T", "O", "Z", "Z", "Z", "T", "I", "L", "O", "S", "O", "I", "T", "S", "O", "O", "I", "S", "Z", "T", "I", "L", "O", "T", "S", "I", "T", "S", "I", "S", "I", "S", "S", "S", "J", "O", "I", "I", "Z", "J", "Z", "J", "S", "Z", "L", "J", "S", "I", "I", "J", "L", "I", "I", "Z", "O", "L", "L", "J", "L", "Z", "I", "O", "T", "L", "S", "J", "Z", "L", "L", "L", "J", "T", "J", "L", "I", "O", "L", "L", "L", "O", "I", "J", "I", "Z", "O", "O", "Z", "T", "I", "S", "O", "Z", "J", "J", "T", "S", "O", "T", "L", "O", "L", "T", "J", "I", "O", "T", "J", "T", "L", "J", "J", "Z", "O", "L", "S", "O", "T", "I", "I", "I", "Z", "L", "T", "L", "S", "I", "J", "T", "S", "L", "T", "L", "I", "S", "Z", "I", "O", "S", "T", "J", "J", "Z", "O", "J", "Z", "T", "J", "J", "L", "I", "L", "I", "O", "L", "J", "J", "T", "S", "S", "Z", "J", "I", "I", "L", "J", "T", "S", "Z", "L", "T", "Z", "O", "T", "Z", "O", "O", "I", "S", "O", "Z", "I", "I", "L", "L", "J", "J", "O", "O", "S", "J", "S", "I", "I", "L", "J", "I", "S", "Z", "I", "I", "T", "Z", "I", "T", "O", "J", "O", "J", "I", "L", "I", "I", "S", "O", "S", "S", "L", "S", "I", "Z", "L", "L", "O", "T", "J", "T", "J", "T", "S", "Z", "J", "L", "I", "I", "S", "T", "Z", "O", "I", "T", "J", "T", "Z", "O", "L", "T", "J", "Z", "O", "S", "I", "J", "T", "O", "L", "S", "Z", "J", "O", "T", "T", "O", "L", "T", "O", "S", "Z", "Z", "Z", "I", "S", "O", "T", "S", "L", "S", "T"}
var g3 = [300]string{"L", "O", "I", "L", "L", "Z", "S", "I", "I", "J", "L", "S", "S", "Z", "T", "S", "I", "L", "J", "Z", "T", "O", "L", "O", "L", "L", "Z", "J", "I", "L", "Z", "S", "T", "T", "J", "L", "Z", "J", "L", "L", "T", "O", "Z", "I", "T", "T", "L", "I", "Z", "J", "L", "Z", "S", "L", "S", "J", "T", "S", "I", "O", "Z", "T", "J", "S", "T", "T", "I", "T", "Z", "O", "T", "S", "I", "Z", "T", "L", "J", "S", "I", "S", "I", "I", "O", "T", "J", "L", "L", "I", "S", "L", "T", "T", "O", "T", "I", "J", "I", "O", "L", "I", "L", "Z", "J", "S", "J", "S", "O", "J", "I", "L", "Z", "Z", "L", "T", "Z", "S", "Z", "Z", "O", "S", "I", "I", "L", "J", "T", "T", "O", "S", "J", "L", "Z", "Z", "I", "I", "T", "J", "O", "J", "Z", "L", "O", "T", "L", "T", "S", "T", "O", "Z", "L", "I", "O", "L", "S", "L", "Z", "J", "O", "L", "O", "J", "S", "L", "I", "T", "Z", "T", "O", "L", "T", "I", "Z", "S", "O", "J", "O", "T", "O", "T", "Z", "Z", "Z", "J", "L", "I", "O", "Z", "L", "S", "T", "L", "J", "S", "L", "S", "Z", "T", "Z", "J", "L", "I", "O", "Z", "S", "J", "I", "J", "T", "L", "J", "Z", "I", "S", "O", "S", "S", "O", "O", "Z", "T", "I", "T", "I", "J", "T", "I", "Z", "O", "O", "Z", "J", "T", "J", "L", "T", "S", "S", "O", "T", "Z", "I", "T", "Z", "L", "Z", "J", "I", "I", "T", "I", "T", "L", "S", "J", "S", "O", "S", "J", "S", "I", "Z", "Z", "I", "T", "T", "S", "T", "J", "J", "O", "I", "L", "S", "J", "I", "L", "Z", "L", "S", "Z", "O", "I", "T", "J", "T", "L", "Z", "Z", "I", "O", "O", "J", "Z", "Z", "I", "S", "T", "T", "J", "S", "Z"}
var g4 = [300]string{"T", "T", "T", "S", "T", "Z", "O", "J", "T", "L", "I", "J", "O", "S", "S", "L", "S", "J", "T", "L", "O", "Z", "T", "I", "T", "I", "S", "Z", "L", "Z", "Z", "J", "Z", "T", "I", "J", "S", "Z", "L", "L", "S", "T", "O", "I", "S", "J", "I", "I", "O", "I", "S", "S", "I", "Z", "S", "O", "O", "L", "O", "L", "J", "J", "S", "I", "S", "I", "O", "I", "Z", "Z", "L", "Z", "Z", "Z", "O", "O", "I", "Z", "I", "O", "I", "I", "Z", "L", "L", "J", "Z", "O", "S", "O", "Z", "Z", "S", "Z", "O", "Z", "I", "J", "S", "S", "L", "S", "T", "J", "S", "O", "T", "J", "Z", "Z", "S", "O", "L", "J", "Z", "L", "L", "J", "O", "I", "L", "J", "L", "Z", "J", "I", "L", "L", "L", "O", "I", "Z", "L", "I", "O", "S", "L", "I", "L", "S", "Z", "Z", "T", "O", "Z", "S", "L", "J", "O", "L", "T", "J", "T", "T", "S", "S", "Z", "J", "T", "J", "L", "I", "L", "O", "L", "Z", "L", "S", "Z", "S", "I", "Z", "Z", "L", "L", "Z", "Z", "S", "Z", "I", "J", "T", "L", "J", "J", "I", "S", "O", "J", "L", "J", "S", "T", "Z", "Z", "J", "L", "O", "J", "I", "Z", "O", "S", "I", "I", "J", "L", "S", "J", "S", "Z", "J", "O", "I", "J", "O", "T", "Z", "L", "I", "J", "J", "I", "S", "J", "Z", "T", "J", "Z", "T", "S", "J", "S", "Z", "T", "Z", "Z", "J", "T", "I", "S", "O", "S", "T", "I", "T", "J", "T", "O", "T", "S", "I", "Z", "Z", "L", "Z", "I", "O", "Z", "S", "O", "T", "L", "L", "S", "I", "Z", "S", "J", "I", "I", "L", "O", "S", "S", "Z", "O", "S", "Z", "S", "O", "S", "O", "J", "J", "L", "T", "L", "J", "S", "J", "L", "O", "S", "O", "Z", "I", "L", "I", "T"}
var g5 = [300]string{"Z", "I", "L", "O", "Z", "L", "L", "J", "S", "T", "J", "I", "Z", "Z", "J", "I", "I", "T", "I", "J", "S", "T", "J", "T", "T", "Z", "L", "S", "O", "S", "I", "L", "S", "L", "T", "L", "T", "S", "I", "Z", "J", "O", "Z", "S", "L", "O", "J", "I", "Z", "L", "O", "T", "S", "Z", "O", "S", "O", "L", "L", "S", "O", "I", "Z", "I", "T", "S", "O", "S", "Z", "S", "J", "J", "Z", "Z", "S", "S", "L", "J", "S", "T", "S", "Z", "Z", "T", "I", "T", "O", "I", "O", "Z", "Z", "I", "O", "L", "L", "J", "Z", "Z", "S", "Z", "O", "S", "L", "J", "J", "I", "O", "J", "Z", "Z", "Z", "I", "S", "J", "T", "Z", "L", "L", "J", "S", "I", "S", "J", "O", "L", "O", "J", "Z", "J", "T", "L", "T", "L", "J", "S", "Z", "Z", "O", "S", "I", "L", "L", "Z", "Z", "S", "I", "I", "L", "S", "O", "T", "Z", "J", "Z", "J", "O", "T", "L", "O", "I", "I", "S", "L", "O", "I", "Z", "T", "I", "I", "S", "O", "L", "L", "I", "I", "S", "O", "I", "S", "I", "O", "O", "O", "I", "S", "J", "Z", "I", "T", "L", "L", "J", "J", "J", "S", "O", "T", "J", "I", "J", "I", "J", "L", "I", "Z", "J", "S", "Z", "S", "T", "Z", "L", "Z", "I", "I", "S", "I", "Z", "S", "T", "O", "L", "T", "Z", "L", "J", "L", "T", "T", "O", "T", "S", "T", "J", "L", "J", "Z", "T", "Z", "J", "Z", "I", "I", "S", "J", "I", "T", "L", "J", "I", "Z", "S", "L", "O", "S", "Z", "I", "O", "S", "Z", "O", "I", "S", "O", "Z", "J", "Z", "J", "O", "Z", "O", "T", "T", "T", "L", "L", "S", "Z", "L", "S", "O", "O", "J", "Z", "T", "L", "Z", "J", "O", "J", "S", "J", "Z", "J", "S", "T", "O", "J", "S", "L"}
var g6 = [300]string{"I", "L", "J", "J", "T", "L", "I", "O", "I", "O", "O", "O", "S", "J", "I", "O", "I", "L", "S", "O", "J", "Z", "T", "S", "Z", "Z", "Z", "T", "I", "I", "T", "J", "J", "S", "S", "L", "S", "O", "J", "I", "O", "S", "Z", "T", "S", "L", "O", "L", "Z", "J", "I", "S", "J", "I", "S", "I", "O", "I", "I", "S", "Z", "Z", "L", "S", "L", "J", "O", "I", "J", "O", "J", "S", "Z", "Z", "Z", "T", "I", "L", "T", "T", "T", "J", "J", "O", "T", "O", "I", "L", "T", "Z", "Z", "I", "Z", "J", "S", "O", "S", "J", "S", "S", "T", "J", "J", "O", "J", "T", "I", "T", "L", "S", "J", "J", "S", "Z", "O", "I", "L", "L", "S", "Z", "I", "J", "O", "Z", "I", "L", "J", "S", "I", "J", "T", "S", "S", "T", "T", "T", "S", "Z", "O", "L", "S", "L", "S", "S", "J", "I", "Z", "Z", "O", "S", "J", "J", "T", "O", "L", "T", "T", "J", "J", "S", "O", "L", "T", "J", "O", "Z", "O", "Z", "Z", "J", "O", "J", "O", "J", "J", "L", "T", "O", "L", "O", "L", "O", "J", "O", "Z", "I", "T", "T", "O", "J", "Z", "S", "O", "L", "L", "I", "T", "O", "I", "T", "Z", "L", "O", "T", "T", "I", "T", "J", "L", "T", "I", "T", "O", "I", "J", "J", "S", "J", "J", "I", "I", "L", "T", "L", "L", "L", "S", "Z", "J", "J", "O", "L", "Z", "S", "T", "Z", "T", "Z", "Z", "I", "O", "T", "T", "S", "O", "T", "L", "L", "O", "O", "Z", "I", "S", "O", "Z", "Z", "O", "I", "J", "L", "L", "O", "I", "O", "J", "O", "Z", "O", "J", "Z", "T", "T", "O", "T", "I", "S", "Z", "Z", "L", "L", "L", "O", "Z", "Z", "O", "O", "T", "S", "T", "I", "Z", "Z", "S", "J", "I", "L", "L", "Z", "J", "I"}
var g7 = [300]string{"S", "O", "O", "J", "J", "J", "L", "Z", "S", "O", "I", "J", "O", "Z", "T", "S", "J", "S", "O", "Z", "T", "J", "J", "J", "I", "T", "O", "J", "T", "I", "S", "S", "I", "J", "I", "I", "I", "T", "I", "L", "I", "I", "S", "L", "T", "Z", "I", "J", "I", "O", "S", "T", "O", "Z", "I", "L", "I", "J", "I", "O", "S", "Z", "L", "S", "I", "I", "Z", "O", "J", "J", "L", "Z", "Z", "O", "I", "I", "I", "L", "I", "L", "S", "J", "L", "J", "O", "Z", "L", "O", "S", "Z", "I", "L", "T", "O", "Z", "S", "J", "S", "I", "S", "J", "T", "Z", "S", "I", "S", "L", "Z", "Z", "T", "S", "O", "S", "O", "T", "J", "S", "O", "J", "S", "O", "S", "I", "Z", "J", "L", "O", "I", "O", "T", "O", "I", "I", "Z", "Z", "T", "J", "Z", "J", "Z", "S", "T", "I", "S", "L", "O", "S", "L", "O", "O", "L", "Z", "J", "J", "S", "J", "S", "I", "O", "S", "S", "J", "S", "S", "T", "J", "L", "L", "L", "J", "Z", "J", "I", "L", "J", "Z", "S", "I", "S", "T", "S", "L", "O", "T", "O", "J", "I", "T", "S", "I", "J", "O", "L", "T", "J", "S", "L", "Z", "Z", "O", "O", "O", "T", "J", "S", "L", "O", "Z", "Z", "J", "J", "I", "L", "S", "Z", "J", "L", "S", "O", "L", "T", "Z", "Z", "S", "Z", "I", "S", "L", "L", "O", "I", "T", "L", "L", "S", "I", "T", "O", "I", "I", "Z", "S", "Z", "J", "L", "S", "T", "S", "Z", "I", "I", "Z", "J", "J", "J", "J", "J", "L", "S", "T", "L", "L", "Z", "O", "J", "T", "O", "S", "T", "J", "L", "J", "L", "T", "J", "I", "S", "O", "S", "S", "T", "I", "T", "O", "O", "J", "L", "Z", "T", "I", "T", "S", "L", "T", "I", "Z", "Z", "L", "O", "J"}
var g8 = [300]string{"O", "O", "O", "L", "Z", "L", "S", "Z", "I", "I", "J", "T", "L", "S", "J", "O", "L", "Z", "J", "I", "L", "I", "I", "S", "J", "L", "L", "I", "O", "O", "I", "S", "Z", "I", "J", "L", "O", "J", "S", "I", "O", "L", "S", "J", "T", "T", "I", "J", "T", "I", "O", "J", "L", "S", "I", "T", "T", "I", "L", "O", "J", "I", "S", "J", "S", "J", "L", "Z", "L", "J", "S", "L", "J", "J", "Z", "J", "S", "T", "O", "O", "O", "S", "Z", "T", "L", "J", "S", "L", "T", "O", "Z", "L", "L", "T", "J", "O", "T", "Z", "Z", "O", "L", "T", "J", "T", "I", "Z", "S", "T", "L", "I", "T", "Z", "I", "I", "L", "T", "T", "I", "J", "T", "O", "Z", "S", "L", "Z", "T", "I", "T", "L", "L", "Z", "L", "T", "I", "Z", "T", "I", "Z", "T", "T", "J", "Z", "O", "S", "Z", "L", "J", "J", "J", "I", "O", "O", "J", "J", "L", "J", "O", "O", "J", "I", "T", "I", "Z", "J", "L", "L", "I", "S", "J", "J", "O", "J", "S", "T", "Z", "J", "O", "S", "L", "I", "I", "S", "O", "S", "I", "T", "J", "O", "O", "J", "O", "L", "S", "I", "I", "L", "Z", "O", "Z", "O", "J", "S", "I", "T", "I", "T", "O", "O", "Z", "Z", "I", "T", "J", "T", "Z", "L", "S", "J", "O", "O", "L", "S", "S", "J", "S", "L", "Z", "L", "J", "T", "O", "J", "O", "L", "O", "Z", "T", "T", "Z", "L", "I", "S", "J", "O", "S", "O", "I", "L", "L", "Z", "I", "O", "O", "L", "O", "S", "T", "O", "L", "T", "S", "S", "O", "J", "I", "O", "Z", "I", "I", "T", "T", "O", "S", "S", "T", "O", "Z", "O", "S", "J", "Z", "T", "O", "O", "J", "L", "O", "O", "S", "J", "T", "T", "O", "L", "J", "O", "S", "O", "S", "O"}
var g9 = [300]string{"J", "L", "Z", "O", "J", "Z", "S", "J", "L", "T", "L", "S", "Z", "L", "S", "T", "S", "S", "L", "I", "J", "T", "S", "Z", "O", "T", "T", "J", "T", "I", "L", "Z", "S", "J", "O", "L", "Z", "I", "O", "Z", "S", "L", "O", "S", "Z", "T", "T", "O", "Z", "Z", "J", "Z", "S", "O", "Z", "L", "L", "I", "T", "S", "J", "L", "I", "J", "J", "T", "J", "I", "O", "L", "I", "S", "L", "I", "O", "S", "J", "T", "I", "L", "Z", "J", "I", "O", "S", "O", "Z", "L", "S", "O", "O", "T", "L", "T", "O", "T", "Z", "O", "Z", "I", "S", "I", "T", "L", "J", "T", "J", "S", "S", "O", "Z", "T", "L", "Z", "I", "Z", "S", "O", "Z", "J", "O", "O", "I", "T", "Z", "I", "S", "S", "T", "I", "O", "S", "Z", "Z", "I", "I", "S", "S", "S", "Z", "Z", "Z", "O", "I", "T", "S", "I", "T", "S", "L", "Z", "S", "S", "O", "S", "O", "S", "S", "Z", "Z", "O", "L", "L", "L", "S", "L", "O", "I", "Z", "J", "T", "S", "J", "T", "O", "I", "S", "S", "L", "O", "I", "Z", "Z", "Z", "J", "J", "J", "O", "J", "Z", "T", "O", "O", "T", "J", "O", "J", "S", "Z", "I", "J", "S", "Z", "Z", "I", "O", "L", "J", "T", "J", "T", "O", "O", "Z", "T", "S", "T", "S", "Z", "J", "I", "S", "I", "O", "T", "J", "I", "I", "I", "O", "T", "I", "Z", "J", "I", "I", "I", "O", "T", "L", "O", "I", "L", "S", "S", "T", "Z", "O", "S", "T", "L", "O", "O", "J", "L", "T", "J", "T", "Z", "O", "J", "T", "T", "T", "Z", "L", "T", "S", "O", "T", "I", "I", "L", "J", "Z", "Z", "S", "J", "I", "J", "S", "L", "J", "I", "I", "I", "J", "J", "S", "L", "L", "I", "J", "L", "T", "I", "J", "J", "J", "T"}
var g10 = [300]string{"T", "O", "I", "J", "S", "I", "O", "Z", "Z", "L", "J", "S", "T", "O", "J", "S", "Z", "Z", "I", "L", "L", "O", "L", "L", "I", "L", "O", "O", "S", "J", "T", "O", "O", "S", "T", "O", "Z", "Z", "S", "I", "J", "L", "Z", "S", "T", "O", "I", "S", "Z", "S", "J", "T", "T", "L", "T", "L", "L", "T", "S", "T", "I", "L", "L", "J", "I", "L", "J", "L", "I", "S", "Z", "I", "Z", "T", "L", "I", "S", "I", "J", "J", "S", "S", "Z", "T", "O", "I", "J", "Z", "T", "J", "J", "Z", "T", "O", "J", "S", "L", "I", "Z", "I", "S", "S", "O", "I", "O", "O", "S", "S", "J", "O", "T", "T", "J", "S", "Z", "T", "J", "T", "O", "O", "L", "Z", "J", "S", "I", "I", "I", "T", "S", "L", "T", "Z", "I", "T", "Z", "Z", "T", "T", "L", "S", "O", "S", "T", "J", "J", "J", "I", "Z", "O", "I", "Z", "I", "T", "L", "L", "L", "T", "J", "L", "L", "S", "L", "O", "S", "S", "Z", "Z", "L", "T", "L", "I", "J", "J", "T", "J", "T", "S", "T", "T", "L", "T", "I", "L", "S", "I", "J", "J", "I", "L", "O", "S", "O", "O", "J", "T", "Z", "I", "O", "L", "O", "Z", "L", "Z", "O", "O", "S", "S", "L", "T", "J", "T", "T", "I", "S", "S", "O", "O", "S", "Z", "T", "O", "T", "S", "J", "I", "J", "I", "T", "S", "T", "I", "L", "T", "T", "S", "J", "O", "Z", "L", "Z", "J", "S", "L", "I", "Z", "J", "T", "L", "O", "O", "T", "Z", "T", "S", "S", "S", "S", "J", "T", "I", "L", "O", "I", "S", "L", "I", "L", "O", "I", "T", "S", "J", "T", "O", "S", "Z", "S", "J", "O", "Z", "T", "J", "T", "J", "I", "J", "L", "Z", "O", "L", "I", "Z", "Z", "J", "S", "T", "S", "I", "O", "Z"}
var g11 = [300]string{"L", "J", "T", "T", "T", "Z", "Z", "T", "I", "S", "J", "L", "T", "O", "I", "Z", "O", "O", "T", "T", "J", "T", "T", "S", "J", "J", "L", "S", "T", "O", "O", "T", "I", "S", "T", "S", "T", "T", "T", "T", "J", "J", "J", "O", "J", "J", "J", "J", "Z", "I", "J", "T", "Z", "T", "I", "S", "S", "Z", "Z", "L", "L", "S", "J", "L", "Z", "I", "I", "J", "S", "T", "S", "Z", "L", "L", "T", "L", "Z", "L", "T", "O", "L", "I", "T", "Z", "I", "S", "S", "J", "Z", "L", "S", "T", "S", "O", "J", "Z", "J", "O", "O", "T", "Z", "S", "T", "Z", "J", "Z", "Z", "I", "O", "L", "S", "Z", "Z", "Z", "T", "T", "O", "L", "O", "O", "L", "Z", "L", "T", "Z", "J", "T", "S", "Z", "Z", "S", "O", "I", "S", "J", "J", "T", "O", "I", "Z", "T", "J", "J", "O", "J", "S", "S", "O", "S", "J", "I", "I", "Z", "I", "J", "J", "Z", "T", "I", "T", "S", "O", "I", "T", "T", "S", "T", "O", "O", "S", "J", "S", "Z", "Z", "S", "I", "L", "O", "O", "I", "T", "L", "I", "T", "Z", "L", "L", "T", "O", "J", "I", "O", "I", "J", "I", "O", "Z", "I", "Z", "I", "S", "J", "L", "L", "L", "S", "T", "I", "O", "T", "S", "L", "O", "J", "L", "O", "L", "L", "J", "S", "T", "I", "O", "J", "I", "J", "J", "J", "L", "L", "I", "J", "L", "S", "L", "S", "L", "J", "Z", "I", "I", "O", "J", "T", "Z", "T", "Z", "Z", "L", "L", "Z", "J", "Z", "J", "T", "O", "S", "J", "O", "J", "L", "Z", "S", "S", "O", "L", "J", "I", "S", "O", "O", "J", "Z", "Z", "Z", "J", "I", "I", "J", "L", "T", "J", "S", "S", "Z", "O", "J", "J", "L", "Z", "T", "J", "O", "Z", "S", "Z", "O", "S", "I", "J"}
var g12 = [300]string{"J", "Z", "L", "O", "L", "T", "O", "S", "S", "O", "T", "S", "L", "I", "J", "S", "O", "J", "O", "I", "O", "L", "L", "Z", "I", "L", "T", "I", "Z", "L", "S", "S", "T", "T", "Z", "L", "T", "Z", "I", "T", "L", "J", "I", "L", "S", "J", "S", "O", "T", "O", "I", "O", "J", "S", "Z", "I", "O", "O", "S", "J", "L", "L", "S", "J", "Z", "O", "T", "Z", "J", "J", "I", "S", "S", "I", "L", "L", "O", "I", "O", "Z", "L", "L", "I", "O", "O", "Z", "Z", "J", "J", "T", "S", "J", "L", "J", "S", "J", "J", "T", "L", "O", "J", "T", "I", "L", "T", "T", "Z", "I", "O", "J", "I", "I", "S", "T", "J", "O", "I", "S", "T", "T", "O", "O", "O", "O", "S", "S", "J", "S", "J", "L", "S", "Z", "J", "T", "L", "J", "O", "S", "L", "J", "O", "L", "J", "L", "T", "I", "J", "L", "L", "L", "L", "O", "J", "J", "L", "S", "Z", "T", "S", "J", "Z", "S", "O", "Z", "J", "Z", "T", "O", "T", "S", "L", "I", "L", "T", "S", "L", "T", "T", "J", "O", "T", "I", "S", "O", "S", "S", "J", "Z", "Z", "O", "J", "S", "Z", "I", "L", "J", "T", "Z", "O", "T", "O", "I", "L", "S", "Z", "S", "I", "J", "Z", "O", "Z", "Z", "T", "T", "L", "I", "T", "L", "I", "J", "J", "T", "Z", "Z", "L", "S", "T", "T", "O", "L", "S", "S", "L", "O", "J", "O", "O", "O", "O", "I", "O", "I", "I", "J", "T", "O", "T", "S", "S", "S", "S", "Z", "T", "S", "T", "L", "O", "O", "L", "L", "J", "O", "Z", "J", "J", "L", "S", "T", "Z", "I", "S", "L", "Z", "S", "T", "J", "I", "S", "I", "O", "I", "S", "L", "I", "S", "O", "I", "J", "I", "O", "L", "J", "S", "J", "L", "I", "L", "L", "T", "S"}
var g13 = [300]string{"I", "I", "T", "O", "O", "O", "L", "Z", "S", "O", "S", "S", "Z", "S", "J", "T", "I", "O", "Z", "J", "S", "Z", "Z", "L", "I", "S", "J", "I", "T", "I", "O", "I", "Z", "O", "O", "Z", "Z", "L", "S", "S", "L", "L", "L", "J", "I", "I", "O", "O", "I", "L", "I", "J", "Z", "J", "J", "J", "Z", "T", "Z", "Z", "O", "L", "O", "J", "J", "T", "S", "S", "T", "Z", "T", "S", "L", "I", "T", "Z", "J", "J", "S", "L", "Z", "I", "J", "O", "L", "J", "J", "I", "L", "L", "Z", "J", "J", "J", "L", "L", "O", "J", "T", "S", "S", "I", "Z", "T", "I", "S", "L", "Z", "J", "I", "Z", "J", "L", "T", "J", "I", "T", "Z", "J", "J", "S", "S", "L", "O", "I", "I", "T", "J", "S", "T", "O", "T", "I", "S", "S", "J", "I", "S", "J", "S", "S", "T", "T", "L", "J", "J", "L", "Z", "S", "L", "S", "Z", "T", "O", "J", "J", "Z", "Z", "T", "Z", "S", "J", "O", "T", "Z", "L", "O", "L", "J", "I", "T", "T", "O", "S", "I", "J", "T", "I", "T", "I", "L", "J", "I", "T", "Z", "J", "Z", "S", "Z", "O", "S", "Z", "Z", "O", "L", "T", "T", "S", "T", "L", "T", "Z", "I", "J", "S", "J", "L", "O", "J", "O", "L", "L", "O", "J", "O", "L", "I", "O", "I", "L", "S", "S", "T", "O", "Z", "I", "J", "S", "I", "S", "I", "O", "L", "O", "J", "I", "T", "Z", "L", "Z", "Z", "T", "Z", "J", "L", "T", "L", "Z", "Z", "O", "O", "S", "I", "S", "L", "J", "Z", "O", "T", "J", "I", "S", "J", "I", "L", "J", "J", "O", "T", "S", "I", "O", "I", "O", "J", "S", "I", "O", "I", "Z", "O", "I", "S", "I", "Z", "I", "S", "O", "L", "S", "Z", "L", "J", "S", "S", "Z", "O", "T", "O", "S"}
var g14 = [300]string{"O", "O", "S", "Z", "L", "O", "O", "O", "I", "L", "J", "J", "S", "S", "J", "I", "L", "S", "O", "Z", "Z", "Z", "I", "T", "L", "J", "O", "L", "J", "Z", "L", "Z", "T", "Z", "S", "L", "T", "L", "T", "J", "L", "T", "T", "L", "S", "I", "J", "S", "S", "J", "I", "Z", "S", "T", "L", "I", "S", "T", "Z", "O", "J", "L", "T", "Z", "O", "T", "L", "L", "I", "J", "I", "T", "S", "L", "O", "I", "T", "T", "O", "T", "I", "I", "S", "Z", "S", "Z", "I", "O", "J", "I", "I", "J", "S", "I", "Z", "S", "O", "T", "J", "S", "S", "L", "L", "S", "I", "S", "J", "O", "T", "T", "L", "O", "T", "Z", "T", "I", "Z", "Z", "T", "S", "T", "L", "J", "S", "I", "L", "J", "S", "Z", "I", "O", "O", "O", "L", "I", "O", "T", "Z", "L", "S", "L", "I", "J", "Z", "L", "J", "T", "L", "T", "S", "Z", "T", "S", "T", "Z", "L", "T", "L", "I", "I", "I", "J", "O", "O", "T", "L", "I", "I", "S", "J", "I", "L", "Z", "Z", "Z", "J", "T", "L", "L", "Z", "J", "O", "Z", "I", "L", "S", "I", "T", "J", "T", "Z", "Z", "I", "Z", "Z", "T", "L", "I", "I", "I", "Z", "O", "I", "L", "T", "S", "I", "I", "T", "Z", "Z", "O", "Z", "S", "Z", "J", "O", "S", "T", "I", "O", "Z", "T", "T", "Z", "S", "T", "L", "Z", "L", "T", "T", "J", "J", "S", "Z", "J", "L", "J", "T", "Z", "Z", "S", "I", "L", "T", "O", "I", "I", "Z", "T", "I", "L", "T", "I", "J", "J", "Z", "J", "Z", "L", "I", "Z", "Z", "I", "J", "I", "S", "O", "O", "Z", "J", "O", "T", "O", "L", "L", "O", "Z", "J", "O", "T", "L", "O", "O", "Z", "Z", "T", "J", "J", "T", "J", "J", "L", "T", "O", "O", "S", "I", "T"}
var g15 = [300]string{"T", "T", "S", "T", "J", "S", "J", "O", "I", "J", "I", "T", "S", "Z", "L", "J", "I", "S", "O", "Z", "Z", "L", "O", "O", "L", "T", "Z", "O", "I", "O", "L", "T", "T", "S", "L", "I", "O", "I", "S", "T", "L", "O", "O", "S", "S", "S", "J", "J", "T", "Z", "Z", "I", "L", "T", "Z", "Z", "J", "Z", "T", "O", "Z", "S", "S", "Z", "T", "Z", "L", "L", "S", "L", "S", "I", "I", "Z", "S", "L", "O", "S", "L", "J", "I", "J", "Z", "J", "S", "L", "T", "O", "L", "I", "Z", "S", "I", "O", "O", "O", "Z", "S", "J", "S", "I", "S", "O", "J", "I", "I", "I", "T", "Z", "T", "Z", "S", "T", "L", "T", "J", "O", "I", "O", "L", "Z", "T", "S", "T", "I", "S", "S", "J", "S", "L", "I", "I", "S", "O", "O", "Z", "Z", "T", "L", "L", "T", "S", "L", "J", "I", "Z", "L", "L", "J", "Z", "L", "Z", "O", "I", "T", "T", "S", "I", "S", "I", "J", "I", "Z", "Z", "Z", "T", "T", "J", "S", "T", "S", "J", "I", "J", "S", "I", "L", "S", "Z", "T", "Z", "I", "S", "S", "I", "O", "O", "T", "S", "T", "O", "Z", "L", "L", "T", "L", "L", "S", "Z", "O", "O", "S", "L", "L", "Z", "L", "I", "I", "T", "T", "J", "S", "J", "T", "L", "I", "T", "Z", "O", "J", "I", "J", "T", "O", "S", "Z", "O", "Z", "S", "L", "S", "O", "Z", "S", "Z", "L", "J", "S", "O", "Z", "Z", "L", "L", "I", "I", "I", "I", "Z", "J", "S", "I", "O", "I", "L", "S", "O", "S", "I", "S", "T", "T", "Z", "O", "J", "S", "O", "J", "L", "J", "J", "O", "T", "L", "I", "T", "I", "I", "S", "O", "I", "O", "Z", "J", "S", "Z", "T", "L", "I", "O", "T", "J", "S", "O", "Z", "O", "I", "J", "O", "I", "J"}
var g16 = [300]string{"T", "L", "Z", "T", "J", "Z", "O", "I", "I", "S", "O", "T", "Z", "Z", "O", "S", "L", "T", "J", "T", "O", "O", "Z", "Z", "S", "J", "J", "L", "S", "S", "J", "J", "Z", "T", "I", "O", "T", "S", "J", "Z", "O", "L", "I", "J", "O", "J", "I", "J", "I", "L", "L", "T", "S", "O", "J", "O", "O", "L", "O", "T", "S", "I", "T", "S", "I", "J", "O", "Z", "J", "L", "S", "Z", "T", "O", "S", "T", "S", "T", "J", "T", "I", "Z", "S", "S", "I", "L", "S", "L", "T", "S", "T", "T", "O", "T", "Z", "J", "S", "I", "T", "J", "Z", "I", "O", "Z", "J", "Z", "J", "O", "I", "O", "Z", "S", "T", "T", "J", "L", "Z", "T", "O", "J", "L", "T", "J", "I", "Z", "S", "I", "O", "J", "I", "O", "Z", "L", "Z", "T", "J", "L", "L", "S", "T", "L", "O", "S", "T", "I", "L", "T", "S", "I", "T", "Z", "T", "I", "I", "L", "Z", "T", "L", "I", "L", "J", "L", "S", "O", "Z", "S", "O", "L", "O", "L", "L", "T", "I", "T", "L", "I", "S", "J", "T", "S", "L", "L", "Z", "S", "I", "Z", "J", "O", "Z", "I", "L", "S", "Z", "L", "L", "T", "O", "J", "S", "O", "I", "T", "I", "L", "J", "S", "J", "Z", "S", "T", "O", "O", "O", "I", "Z", "L", "T", "I", "I", "I", "T", "I", "J", "T", "Z", "O", "O", "I", "Z", "O", "J", "O", "J", "T", "J", "I", "J", "L", "J", "L", "J", "J", "O", "Z", "I", "L", "Z", "S", "S", "T", "T", "T", "T", "Z", "O", "I", "I", "T", "T", "J", "I", "J", "L", "I", "T", "Z", "I", "Z", "T", "T", "T", "T", "L", "Z", "L", "O", "J", "T", "O", "T", "Z", "S", "S", "L", "J", "T", "O", "S", "L", "S", "J", "O", "O", "O", "S", "L", "I", "S", "J", "I"}
var g17 = [300]string{"S", "Z", "T", "T", "T", "I", "S", "S", "L", "J", "I", "Z", "I", "S", "O", "Z", "O", "J", "I", "Z", "Z", "J", "Z", "I", "T", "Z", "L", "O", "T", "Z", "T", "I", "Z", "Z", "L", "I", "Z", "Z", "S", "L", "Z", "I", "L", "J", "Z", "I", "J", "T", "O", "Z", "I", "J", "Z", "I", "O", "I", "O", "I", "T", "L", "S", "J", "S", "J", "Z", "Z", "J", "L", "O", "O", "O", "I", "O", "T", "O", "J", "L", "T", "I", "O", "O", "Z", "S", "T", "S", "L", "J", "Z", "I", "T", "O", "Z", "I", "O", "I", "L", "S", "T", "O", "S", "T", "J", "S", "L", "J", "I", "J", "L", "T", "J", "O", "T", "T", "T", "Z", "S", "T", "O", "T", "J", "J", "Z", "J", "J", "Z", "Z", "J", "L", "J", "I", "T", "T", "L", "T", "L", "L", "L", "J", "I", "T", "S", "S", "I", "I", "I", "Z", "Z", "Z", "S", "L", "T", "T", "O", "Z", "O", "T", "L", "T", "O", "I", "Z", "Z", "O", "S", "Z", "O", "L", "J", "I", "I", "I", "Z", "Z", "S", "O", "J", "I", "I", "S", "J", "Z", "I", "I", "I", "L", "S", "J", "S", "L", "I", "S", "T", "S", "Z", "T", "Z", "O", "O", "Z", "S", "O", "S", "O", "T", "S", "T", "Z", "J", "Z", "J", "T", "L", "T", "L", "T", "J", "S", "Z", "O", "O", "T", "S", "L", "I", "T", "Z", "T", "T", "L", "S", "Z", "L", "S", "J", "J", "Z", "O", "T", "I", "O", "Z", "I", "T", "S", "T", "I", "L", "O", "L", "J", "J", "J", "I", "I", "T", "I", "T", "S", "S", "S", "O", "S", "O", "Z", "T", "Z", "S", "S", "L", "S", "T", "S", "S", "J", "O", "J", "O", "L", "O", "S", "Z", "Z", "O", "S", "L", "I", "S", "S", "I", "O", "S", "Z", "L", "T", "O", "Z", "I", "O", "I", "I"}
var g18 = [300]string{"L", "L", "T", "I", "Z", "S", "Z", "Z", "I", "O", "Z", "J", "S", "O", "I", "T", "T", "S", "S", "O", "L", "Z", "J", "T", "S", "I", "T", "T", "S", "Z", "I", "Z", "L", "J", "L", "I", "O", "O", "I", "Z", "L", "L", "L", "O", "Z", "L", "J", "S", "O", "S", "J", "I", "L", "I", "T", "O", "L", "S", "S", "L", "O", "L", "S", "Z", "Z", "Z", "T", "J", "Z", "S", "Z", "O", "S", "O", "I", "Z", "T", "T", "I", "Z", "Z", "T", "T", "Z", "S", "T", "I", "S", "O", "J", "S", "J", "L", "Z", "S", "L", "S", "Z", "Z", "T", "Z", "I", "T", "L", "I", "I", "Z", "T", "I", "Z", "J", "J", "I", "L", "I", "L", "I", "O", "O", "J", "S", "L", "I", "J", "L", "J", "L", "S", "O", "Z", "L", "I", "L", "Z", "I", "Z", "J", "I", "I", "J", "T", "Z", "S", "I", "O", "O", "T", "Z", "L", "S", "J", "S", "S", "I", "L", "L", "O", "J", "J", "Z", "T", "L", "I", "T", "L", "I", "T", "I", "J", "O", "Z", "Z", "L", "J", "O", "O", "Z", "T", "S", "L", "S", "T", "T", "T", "I", "I", "S", "Z", "O", "Z", "S", "S", "S", "S", "S", "L", "O", "O", "S", "Z", "Z", "Z", "L", "Z", "Z", "Z", "J", "I", "Z", "O", "Z", "S", "Z", "J", "S", "O", "T", "T", "T", "S", "I", "O", "L", "O", "J", "S", "J", "I", "S", "Z", "T", "S", "O", "I", "I", "T", "O", "Z", "T", "I", "I", "J", "L", "O", "O", "S", "L", "L", "Z", "Z", "Z", "O", "L", "S", "I", "L", "Z", "Z", "L", "Z", "I", "O", "I", "O", "Z", "Z", "O", "Z", "I", "I", "S", "T", "O", "S", "L", "T", "J", "S", "J", "L", "S", "I", "T", "T", "L", "L", "T", "S", "L", "T", "S", "S", "Z", "S", "J", "L", "J", "T", "Z", "O"}
var g19 = [300]string{"Z", "L", "T", "I", "L", "I", "L", "L", "O", "Z", "Z", "S", "L", "O", "O", "O", "I", "T", "I", "T", "O", "J", "T", "I", "S", "Z", "T", "Z", "L", "O", "O", "S", "T", "L", "Z", "T", "Z", "T", "S", "J", "L", "Z", "I", "O", "Z", "Z", "L", "I", "O", "O", "O", "J", "L", "Z", "J", "I", "S", "L", "I", "O", "S", "I", "L", "I", "S", "I", "I", "O", "O", "J", "L", "L", "I", "S", "S", "I", "S", "S", "L", "L", "O", "Z", "O", "Z", "T", "I", "I", "Z", "Z", "S", "L", "Z", "S", "O", "Z", "J", "O", "I", "L", "T", "S", "O", "S", "S", "O", "S", "T", "T", "S", "J", "T", "O", "T", "T", "J", "J", "O", "Z", "S", "O", "J", "O", "O", "T", "I", "L", "J", "L", "O", "O", "J", "Z", "I", "I", "T", "J", "T", "S", "J", "T", "O", "S", "S", "S", "T", "J", "S", "L", "S", "T", "I", "O", "O", "J", "O", "J", "I", "T", "T", "S", "J", "L", "S", "L", "S", "O", "T", "Z", "Z", "J", "O", "O", "Z", "T", "J", "L", "O", "O", "O", "L", "Z", "S", "J", "S", "Z", "I", "L", "J", "T", "J", "T", "J", "J", "O", "I", "L", "Z", "J", "Z", "J", "Z", "J", "L", "L", "T", "O", "I", "Z", "L", "T", "S", "L", "T", "O", "O", "S", "J", "O", "Z", "J", "L", "I", "S", "L", "Z", "Z", "Z", "Z", "S", "L", "I", "T", "J", "S", "L", "S", "O", "L", "I", "O", "T", "I", "O", "T", "T", "S", "I", "I", "O", "T", "S", "J", "I", "T", "J", "J", "S", "L", "T", "L", "J", "O", "S", "T", "S", "I", "Z", "S", "Z", "L", "L", "L", "J", "O", "L", "S", "J", "Z", "T", "J", "I", "S", "S", "J", "Z", "L", "I", "O", "S", "O", "L", "J", "J", "T", "O", "I", "L", "J", "S", "S"}
var g20 = [300]string{"L", "L", "T", "I", "L", "J", "S", "S", "S", "T", "O", "J", "Z", "Z", "J", "Z", "T", "J", "I", "Z", "L", "I", "O", "T", "L", "T", "S", "I", "J", "I", "I", "I", "T", "Z", "J", "J", "T", "O", "Z", "I", "J", "Z", "S", "O", "O", "T", "O", "J", "L", "Z", "L", "Z", "S", "L", "S", "T", "T", "S", "O", "J", "I", "Z", "Z", "L", "I", "L", "O", "T", "L", "T", "S", "O", "Z", "T", "I", "T", "O", "Z", "J", "S", "T", "J", "S", "S", "I", "L", "Z", "J", "O", "S", "J", "J", "L", "L", "J", "Z", "J", "T", "L", "Z", "Z", "O", "T", "S", "J", "S", "O", "T", "J", "J", "S", "O", "J", "L", "S", "T", "Z", "O", "T", "L", "Z", "J", "T", "Z", "J", "S", "T", "S", "S", "L", "O", "Z", "J", "S", "T", "I", "O", "T", "I", "Z", "I", "J", "T", "S", "O", "J", "L", "I", "T", "Z", "T", "S", "I", "S", "L", "L", "T", "I", "S", "Z", "S", "S", "I", "Z", "I", "T", "L", "L", "S", "Z", "O", "L", "T", "O", "T", "T", "T", "J", "J", "Z", "O", "J", "L", "S", "I", "I", "J", "S", "J", "S", "L", "O", "J", "J", "J", "Z", "Z", "T", "T", "L", "L", "T", "J", "Z", "O", "O", "T", "T", "O", "J", "J", "J", "S", "Z", "T", "Z", "Z", "Z", "J", "I", "Z", "Z", "Z", "Z", "Z", "L", "Z", "I", "L", "O", "S", "J", "S", "L", "T", "I", "Z", "S", "O", "I", "I", "O", "J", "I", "O", "J", "Z", "L", "O", "Z", "T", "J", "I", "S", "Z", "L", "I", "J", "T", "T", "S", "L", "Z", "T", "T", "T", "J", "S", "I", "Z", "S", "O", "J", "S", "S", "I", "J", "I", "Z", "L", "O", "T", "Z", "S", "I", "O", "Z", "J", "I", "T", "J", "O", "T", "J", "I", "L", "Z", "T", "O", "S"}

var defaultStrategy = Strategy{
	Burn:   4,
	BHoles: 14,
	FHoles: 5,
	CHoles: 1,
	HighY:  2,
	Step:   3,
}

/*
func Test_generate(t *testing.T) {
	for j := 1; j <= 20; j++ {

		i := 1
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("var g", j, " =[300]string{\"", pieces[rand.Intn(len(pieces))], "\"")
		for i < 300 {
			fmt.Print(",\"", pieces[rand.Intn(len(pieces))], "\"")
			i++
		}
		fmt.Print("}")
		fmt.Println()

	}

}
*/
func Benchmark_moves(b *testing.B) {
	for n := 0; n < b.N; n++ {
		game := Game{Strategy: gameSt}
		game.asignSettings("timebank", "10000")
		game.asignSettings("time_per_move", "500")
		game.asignSettings("player_names", "player1,player2")
		game.asignSettings("your_bot", "player1")
		game.asignSettings("field_width", "10")
		game.asignSettings("field_height", "20")

		rand.Seed(time.Now().UTC().UnixNano())

		row1 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row1[i] = strconv.Itoa(rand.Intn(3))
		}
		row2 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row2[i] = strconv.Itoa(rand.Intn(3))
		}
		row3 := make([]string, 10, 10)
		for i := 0; i < 10; i++ {
			row3[i] = strconv.Itoa(rand.Intn(3))
		}

		game.asignUpdates("game", "round", "4")
		game.asignUpdates("game", "this_piece_type", pieces[rand.Intn(len(pieces))])
		game.asignUpdates("game", "next_piece_type", pieces[rand.Intn(len(pieces))])
		game.asignUpdates("game", "this_piece_position", "3,-1")
		game.asignUpdates("player1", "field", "0,0,0,1,1,1,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;0,0,0,0,0,0,0,0,0,0;"+strings.Join(row1, ",")+";"+strings.Join(row2, ",")+";"+strings.Join(row3, ","))
		game.asignUpdates("player1", "row_points", "0")
		game.asignUpdates("player1", "combo", "0")
		game.calculateMoves()
	}
}

func Benchmark_fixholes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testField := Field{{false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, true, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}, {false, false, false, false, false, false, false, false, false, false}}
		piece := Piece{Name: "T", Rotation: 0}
		piece.InitSpace(Cell{X: 3, Y: 19})
		hole := Cell{X: 5, Y: 0}
		testField.FixHoles(piece, []Cell{hole}, testField.Picks().Max())
	}
}

func Benchmark_many(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGame(&Game{Strategy: defaultStrategy}, g1, false)
	}
}

func Benchmark_one(b *testing.B) {
	for n := 0; n < b.N; n++ {
		playGame(&Game{Strategy: defaultStrategy}, g1, true)
	}
}

func Benchmark_strategy(banch *testing.B) {
	fmt.Println("")
	fmt.Println("Burn	BHoles	FHoles	CHoles	HighY	Step	Score	minS	maxS	Round	minR	maxR")
	for n := 0; n < banch.N; n++ {
		for b := 1; b <= 5; b++ {
			for bh := 5; bh <= 10; bh++ {
				for fh := 1; fh <= 5; fh++ {
					//for ch := 1; ch <= 1; ch++ {
					for hy := 1; hy <= 3; hy++ {
						for s := 1; s <= 3; s++ {
							//							st := Strategy{Burn: b, BHoles: bh, FHoles: fh, CHoles: 1, HighY: hy, Step: s}
							//go playGames(st, 22, false, false)
						}
					}
					//fmt.Println("start sleep")
					time.Sleep(95000000000)
					//fmt.Println("end sleep")
					//}
					//time.Sleep(30000000000)
				}
				time.Sleep(30000000000)
			}
		}
		//save("strategies", strategies)
	}
}

func playGame(g *Game, input [300]string, visual bool) (int, int) {
	g.asignSettings("player_names", "player1,player2")
	g.asignSettings("your_bot", "player1")
	g.Round = 0
	g.MyPlayer.Points = 0
	g.MyPlayer.Field = initialField
	g.MyPlayer.Picks = initialField.Picks()
	position := &Piece{}
	position.FieldAfter = initialField
	assignPieces(g, input[0])
	keepGoing := true

	i := 0
	for keepGoing {
		applyPoints(g, position)
		position.FieldAfter.Burn()
		g.MyPlayer.Field = position.FieldAfter
		addSolidLines(g)
		addGarbageLines(g)
		g.MyPlayer.Picks = g.MyPlayer.Field.Picks()
		assignPieces(g, input[i])
		g.Round++

		if visual {
			//fmt.Println("D", position.Damage, "S", position.Score)
			fmt.Println(g.CurrentPiece.Name, "sore:", g.MyPlayer.Points, "round:", g.Round, "combo:", g.MyPlayer.Combo)
			if position.Moves == "" {
				fmt.Println("drop")
			} else {
				fmt.Println(position.Moves + ",drop")
			}
			PrintVisual(g.MyPlayer.Field)
			time.Sleep(1000000000)
		}

		position = g.calculateMoves()

		if position == nil ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][3] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][4] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][5] ||
			g.MyPlayer.Field[g.MyPlayer.Field.Height()-1][6] {
			keepGoing = false
		}
		i++
	}

	return g.Round, g.MyPlayer.Points
}

func assignPieces(g *Game, piece string) {
	g.CurrentPiece = g.NextPiece
	x := 3
	if piece == "O" {
		x = 4
	}
	g.NextPiece = Piece{Name: piece, Rotation: 0}
	g.NextPiece.InitSpace(Cell{x, g.MyPlayer.Field.Height() - 1})
}

func applyPoints(g *Game, p *Piece) {
	if g.Round > 1 {
		points := p.getPoints(g.MyPlayer.Combo)
		if points > 0 {
			g.MyPlayer.Combo++
		} else {
			g.MyPlayer.Combo = 0
		}
		g.MyPlayer.Points += points
	}
}

func isRoof(g *Game) bool {
	for _, col := range g.MyPlayer.Field[g.MyPlayer.Field.Height()-1] {
		if col {
			//fmt.Println("roof", g.MyPlayer.Field.Height())
			return true
		}
	}
	return false
}

func addSolidLines(g *Game) {
	r := g.Round % 20
	if r == 0 {
		g.MyPlayer.Field = g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1]
		g.Height = g.Height - 1
	}
}

func addGarbageLines(g *Game) {
	r := g.Round % 5
	if r == 0 && g.Round != 0 {
		size := g.MyPlayer.Field.Width()
		row := make([]bool, size)
		for i := range row {
			row[i] = true
		}
		hole := rand.Intn(size)
		row[hole] = false
		g.MyPlayer.Field = append([][]bool{row}, [][]bool(g.MyPlayer.Field[:g.MyPlayer.Field.Height()-1])...)
	}
}

func save(fileName string, records [][]string) {
	csvfile, err := os.Create("output/" + fileName + ".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}

func statistic(a []int) (int, int, int) {
	if len(a) < 4 {
		return a[0], a[0], a[0]
	}
	sort.Ints(a)
	var total int
	for i := 1; i < len(a)-1; i++ {
		total += a[i]
	}
	avr := total/len(a) - 2
	return avr, a[1], a[len(a)-2]
}
