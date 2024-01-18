// Package reusable_const_vars
/*
This file contains USA states and postal codes

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
package reusable_const_vars

//goland:noinspection ALL
const (
	// output Modes
	MODE_OUTPUT_DISPLAY = "display"
	MODE_OUTPUT_LOG     = "log"
	MODE_RELEASE        = "release"
	// debug
	DEBUG_MESSAGE = "DEBUG_MESSAGE:"
	DEBUG_MODE    = "DEBUG"
	// Fully qualified filenames
	FQN_CERTIFICATE          = "CertificateFQN"
	FQN_FIREBASE_CREDENTIALS = "FirebaseCredentialsFQN"
	FQN_GCP_CREDENTIALS      = "GCPCredentialsFQN"
	FQN_NATS_CREDENTIALS     = "NATSCredsFQN"
	FQN_TLS_CABUNDLE         = "TLS:TLSCABundle"
	FQN_TLS_CERTIFICATE      = "TLS:TLSCert"
	FQN_TLS_PRIVATE_KEY      = "TLS:TLSKey"
	// Operators
	OPER_DOUBE_EQUAL_SIGN = "=="
	OPER_EQUAL_SIGN       = "="
	// Text Strings
	TXT_BYPASS                   = "bypass"
	TXT_DIRECTORY_EXISTS         = "Directory exists"
	TXT_DIRECTORY_DOES_NOT_EXIST = "Directory doesn't exist"
	TXT_DIRECTORY_PARAM_EMPTY    = "Directory parameter is empty"
	TXT_DIRECTORY                = "DIRECTORY: "
	TXT_FILENAME                 = "FILENAME:"
	TXT_FILENAME_EXISTS          = "Filename exists"
	TXT_FILENAME_DOES_NOT_EXISTS = "Filename doesn'ts exist"
	TXT_IS_EMPTY                 = "is empty"
	TXT_IS_MISSING               = "is missing"
	TXT_IS_OK                    = "is ok"
	TXT_IS_UNREADABLE            = "is unreadable"
	TXT_EMAIL                    = "EMAIL"
	TXT_EMPTY                    = "empty"
	TXT_EVIRONMENT               = "ENVIRONMENT: "
	TXT_FALSE                    = "false"
	TXT_GOOD                     = "good"
	TXT_GOT_WRONG_BOOLEAN        = "Got wrong boolean value."
	TXT_INTO                     = "into"
	TXT_INVALID                  = "INVALID"
	TXT_LINE_SEPARATOR           = "----------------"
	TXT_MESSAGE                  = "message"
	TXT_MISSING                  = "MISSING"
	TXT_MISSING_KEY              = "missing key"
	TXT_MISSING_VALUE            = "missing value"
	TXT_NA                       = "N.A."
	TXT_NIL                      = "nil"
	TXT_NONE                     = "none"
	TXT_NOT_FOUND                = "not found"
	TXT_PROCESS_MANUALLY         = "Process transfer manually."
	TXT_PROTECTED                = "PROTECTED"
	TXT_OUTOF                    = "out of"
	TXT_PHONE                    = "phone"
	TXT_SERVER_NAME              = "SERVER_NAME: "
	TXT_SERVER_VERSION           = "SERVER VERSION: "
	TXT_TRUE                     = "true"
	// Software
	LOCAL_HOST              = "localhost"
	PID_FILENAME            = "/server.pid"
	DEFAULT_VERSION         = "9999.9999.9999"
	ENVIRONMENT_LOCAL       = "local"
	ENVIRONMENT_DEVELOPMENT = "development"
	ENVIRONMENT_PRODUCTION  = "production"
	// Testing
	TEST_POSITVE_SUCCESS  = "Positive Case: Successful: "
	TEST_NEGATIVE_SUCCESS = "Negative Case: Successful: "
	// Unassigned
	ATTEMPT_LIMIT_EXCEEDED  = "Attempt limit exceeded"
	AUTHENTICATOR_SERVICE   = "AuthenticatorService"
	NATS_NON_TLS_CONNECTION = "NON-TLS"
	NATS_TLS_CONNECTION     = "TLS "
	VALID                   = "VALID"
	// Values
	VAL_EMPTY = ""
	VAL_ZERO  = 0
)
