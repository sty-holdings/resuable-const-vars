/*
This file contains country names and codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
library reusable_const_vars;

const CAN_AB = "AB"; // Alberta
const CAN_BC = "BC"; // British Columbia
const CAN_MB = "MB"; // Manitoba
const CAN_NB = "NB"; // New Brunswick
const CAN_NL = "NL"; // Newfoundland and Labrador
const CAN_NT = "NT"; // Northwest Territories
const CAN_NS = "NS"; // Nova Scotia
const CAN_NU = "NU"; // Nunavut
const CAN_ON = "ON"; // Ontario
const CAN_PE = "PE"; // Prince Edward Island
const CAN_QC = "QC"; // Quebec
const CAN_SK = "SK"; // Saskatchewan
const CAN_YT = "YT"; // Yukon
const CAN_ALL_PROVIDENCE_CODES = "AB BC MB NB NL NT NS NU ON PE QC SK YT";

class CanadaProvinceInfo {
  final String englishName;
  final String frenchName;

  const CanadaProvinceInfo({
    required this.englishName,
    required this.frenchName,
  });
}

const canadianProvinceInfo = {
  'AB': CanadaProvinceInfo(englishName: 'Alberta', frenchName: ''),
  'BC': CanadaProvinceInfo(englishName: 'British Columbia', frenchName: ''),
  'MB': CanadaProvinceInfo(englishName: 'Manitoba', frenchName: ''),
  'NB': CanadaProvinceInfo(englishName: 'New Brunswick', frenchName: ''),
  'NL': CanadaProvinceInfo(
    englishName: 'Newfoundland and Labrador',
    frenchName: '',
  ),
  'NT': CanadaProvinceInfo(englishName: 'Northwest Territories', frenchName: ''),
  'NS': CanadaProvinceInfo(englishName: 'Nova Scotia', frenchName: ''),
  'NU': CanadaProvinceInfo(englishName: 'Nunavut', frenchName: ''),
  'ON': CanadaProvinceInfo(englishName: 'Ontario', frenchName: ''),
  'PE': CanadaProvinceInfo(englishName: 'Prince Edward Island', frenchName: ''),
  'QC': CanadaProvinceInfo(englishName: 'Quebec', frenchName: ''),
  'SK': CanadaProvinceInfo(englishName: 'Saskatchewan', frenchName: ''),
  'YT': CanadaProvinceInfo(englishName: 'Yukon', frenchName: ''),
};

String getEnglishProvinceName(String postalCode) {
  return canadianProvinceInfo[postalCode]!.englishName;
}

String getFrenchProvinceName(String postalCode) {
  return canadianProvinceInfo[postalCode]!.frenchName;
}
