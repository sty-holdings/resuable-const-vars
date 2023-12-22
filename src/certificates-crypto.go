// Package reusable_const_vars
/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

NOTES:
    {Enter any additional notes that you believe will help the next developer.}

COPYRIGHT:
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
package reusable_const_vars

//goland:noinspection ALL
const (
	CERT_ED25519         = "ED25519"
	CERT_RSA             = "RSA"
	CERT_RS265           = "RS256"
	CERT_ECDSACURVE      = "ECDSACURVE"
	CERT_ECDSACURVE_P224 = "P224"
	CERT_ECDSACURVE_P256 = "P256"
	CERT_ECDSACURVE_P384 = "P384"
	CERT_ECDSACURVE_P521 = "P521"
	CERT_PRIVATE_KEY     = "RSA PRIVATE KEY"
	CERT_PUBLIC_KEY      = "PUBLIC KEY"
	CERTIFICATE          = "CERTIFICATE"
)
