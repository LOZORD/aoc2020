package main

import "testing"

/*
INVALID
eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
*/

/*
VALID
pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
*/

func TestIsValid(t *testing.T) {
	*part2Flag = true

	for _, tc := range []struct {
		name      string
		m         map[string]string
		wantValid bool
	}{{
		m: map[string]string{
			"eyr": "1972", "cid": "100",
			"hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926",
		},
		wantValid: false,
	}, {
		m: map[string]string{
			"iyr": "2019",
			"hcl": "#602927", "eyr": "1967", "hgt": "170cm",
			"ecl": "grn", "pid": "012533040", "byr": "1946",
		},
		wantValid: false,
	}, {
		m: map[string]string{
			"hcl": "dab227", "iyr": "2012",
			"ecl": "brn", "hgt": "182cm", "pid": "021572410", "eyr": "2020", "byr": "1992", "cid": "277",
		},
		wantValid: false,
	}, {
		m: map[string]string{
			"hgt": "59cm", "ecl": "zzz",
			"eyr": "2038", "hcl": "74454a", "iyr": "2023",
			"pid": "3556412378", "byr": "2007",
		},
		wantValid: false,
	}, {
		m: map[string]string{
			"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980",
			"hcl": "#623a2f",
		},
		wantValid: true,
	}, {
		m: map[string]string{
			"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989",
			"iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm",
		},
		wantValid: true,
	}, {
		m: map[string]string{
			"hcl": "#888785",
			"hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88",
			"pid": "545766238", "ecl": "hzl",
			"eyr": "2022",
		},
		wantValid: true,
	}, {
		m: map[string]string{
			"iyr": "2010", "hgt": "158cm", "hcl": "#b6652a", "ecl": "blu", "byr": "1944", "eyr": "2021", "pid": "093154719",
		},
		wantValid: true,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			gotValid := isValid(tc.m)
			if gotValid != tc.wantValid {
				t.Errorf("isValid(%v) = %v, want %v", tc.m, gotValid, tc.wantValid)
			}
		})
	}
}

func TestPid(t *testing.T) {
	if !prx.MatchString("000000001") {
		t.Error("000000001 failed, want succeeded")
	}
	if prx.MatchString("0123456789") {
		t.Error("0123456789 succeeded, want failed")
	}
}
