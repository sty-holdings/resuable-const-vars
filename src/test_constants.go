// Package reusable_const_vars
/*
This file contains USA states and postal codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.

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
	// Must be updated for testing
	TEST_PLAID_PUBLIC_ACCESS_TOKEN  = "public-sandbox-01894116-125c-4901-bcd1-9e6206a4f72b" // This needs to be update by running Plaid quickstart (https://plaid.com/docs/quickstart/)
	TEST_CITIZEN_PLAID_ACCESS_TOKEN = "access-sandbox-c520b7ab-cfba-485c-b75e-e8102f1c7b9f" // This is for Citizen Bank - DO NOT CHANGE
	TEST_CITIZEN_PLAID_ITEM_ID      = "qP46VRGLA9H9lwDMaE5qcnGnzxVg8zCdBjwxr"               // This is for Citizen Bank - DO NOT CHANGE
	TEST_CITIZEN_PLAID_ACCOUNT_ID   = "wMQ6VKGWqJcl8odjErJ6cz99kGDQZgtE6Dw17"               // This is for Citizen Bank - DO NOT CHANGE
	//
	TEST_REQUESTED_BUNDLE_ALLOCATION   = 1
	TEST_ALLOCATION_MIN                = 25
	TEST_ALLOCATION_MAX                = 9000
	TEST_APPLICATION                   = "TEST_APPLICATION"
	TEST_AWS_INFORMATION_FQN           = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/config/development/AWS-account-info.json"
	TEST_AWS_KEYSET_URL                = "https://cognito-idp.us-west-2.amazonaws.com/%v/.well-known/jwks.json"
	TEST_AWS_RAW_ACCESS_TOKEN_FQN      = "/Users/syacko/workspace/styh-dev/src/albert/core/testTokens/token_raw_access"
	TEST_AWS_RAW_ID_TOKEN_FQN          = "/Users/syacko/workspace/styh-dev/src/albert/core/testTokens/token_raw_id"
	TEST_AWS_TEST_TOKEN_FQN            = "/Users/syacko/workspace/styh-dev/src/albert/core/testTokens/AWSTokens"
	TEST_BUNDLE_ID_0                   = "TEST_BUNDLE_0"
	TEST_BUNDLE_ID_1                   = "TEST_BUNDLE_1"
	TEST_BUNDLE_ID_2                   = "TEST_BUNDLE_2"
	TEST_BUNDLE_ID_F                   = "TEST_BUNDLE_%v"
	TEST_CA_BUNDLE_FQN                 = "/Volumes/development-share/.keys/savup/STAR_savup_com/CAbundle.crt"
	TEST_CERTIFICATE_FQN               = "/Volumes/development-share/.keys/savup/STAR_savup_com/STAR_savup_com.crt"
	TEST_CITY                          = "Mountain House"
	TEST_CONFIGURATION_FQN             = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/config/development/local-savup-config.json"
	TEST_CONFIGURATION_FQN_INVALID     = "/Users/syacko/workspace/styh-dev/src/albert/SavUpServer/testing/files/BAD_savup-development.json"
	TEST_CONFIGURATION_WTIH_TLS_FQN    = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/config/local/local-savup-config-with-TLS.json"
	TEST_DATASTORE                     = "TEST_DATASTORE"
	TEST_DATASTORE_SUBCOLLECTION       = "TEST_DATASTORE_SUBCOLLECTION"
	TEST_DOCUMENT_ID_0                 = "TEST_DOCUMENT_0"
	TEST_DOCUMENT_ID_F                 = "TEST_DOCUMENT_%v"
	TEST_EMAIL_ADDRESS                 = "scott.yacko@sty-holdings.com"
	TEST_EMAIL_NAME                    = "test example"
	TEST_EMAIL_SUBJECT                 = "TEST_SUBJECT"
	TEST_ENVIRONMENT                   = "TEST_ENVIRONMENT"
	TEST_FIELD_NAME                    = "TEST_FIELD_NAME"
	TEST_FIELD_VALUE                   = "TEST_FIELD_VALUE"
	TEST_FIREBASE_CREDENTIALS          = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/firebase-creds.json"
	TEST_FIREBASE_CREDENTIALS_INVALID  = "/Users/syacko/workspace/styh-dev/src/albert/SavUpServer/testing/files/TEST-BAD-FIREBASE-CREDENTIALS.json"
	TEST_GCP_CREDENTIALS               = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/gcp-creds.json"
	TEST_GCP_CREDENTIALS_INVALID       = "/Users/syacko/workspace/styh-dev/src/albert/SavUpServer/testing/files/TEST-BAD-FIREBASE-CREDENTIALS.json"
	TEST_GOOD_FQN                      = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/testing/files/good-file.json"
	TEST_GOOD_PURPOSE                  = "TEST_GOOD_PURPOSE"
	TEST_HOST                          = "TEST_HOST"
	TEST_INSTITUTION_CITIZEN_BANK      = "Citizen"
	TEST_INSTITUTION_CITIZEN_CLONE     = "CitizenClone"
	TEST_INSTITUTION_CHASE             = "Chase"
	TEST_INSTITUTION_CHASE_CLONE       = "ChaseClone"
	TEST_INTEREST_RATE_HIGH            = 10.0
	TEST_INTEREST_RATE_LOW             = 0.4
	TEST_INTEREST_RATE_MED             = 5.2
	TEST_JSON_INVALID                  = `{"name": "Test Nam`
	TEST_MALFORMED_JSON_FILE           = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/testing/files/TEST-MALFORMED-JSON-FILE.json"
	TEST_MESSAGE_LABEL                 = "Test Message"
	TEST_NATS_CREDENTIALS              = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/nats-savup.creds"
	TEST_NATS_SYS_CREDENTIALS          = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/nats-sys.creds"
	TEST_NATS_URL                      = "nats-prod-0001.savup.com"
	TEST_NO_SUCH_DIRECTORY             = "/TEST-NO-SUCH-DIRECTORY.DELETE_ME"
	TEST_NO_SUCH_FILE                  = "/tmp/TEST-NO-SUCH-FILE.DELETE_ME"
	TEST_OFFERING_SIZE                 = 250000
	TEST_PLAID_ACCESS_TOKEN_FQN        = "/Users/syacko/workspace/styh-dev/src/albert/core/testTokens/plaid_access_token"
	TEST_PLAID_ACCOUNT_ID              = "QnzB1vQgZRiGv5jad8WnH9L6vlaonnsG85Apk"
	TEST_PLAID_ITEM_ID_FQN             = "/Users/syacko/workspace/styh-dev/src/albert/core/testValues/plaid_item_id"
	TEST_PLAID_PUBLIC_TOKEN_FQN        = "/Users/syacko/workspace/styh-dev/src/albert/core/testTokens/plaid_public_token"
	TEST_PID_DIRECTORY                 = "/tmp"
	TEST_PRIVATE_KEY_FQN               = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/test_key"
	TEST_REGION                        = "us-west-2"
	TEST_REPLY_INVALID                 = "Bad Reply"
	TEST_SAVUP_PRIVATE_KEY_FQN         = "/Volumes/development-share/.keys/savup/savup.com.key"
	TEST_SENDER_NAME                   = "Support"
	TEST_SENDGRID_KEY_FILE             = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/sendgrid-dev-creds.json"
	TEST_SERVER_PID_FILE               = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/.run/server.pid"
	TEST_STRING                        = "TEST"
	TEST_STRIPE_CUSTOMER_ACCOUNT_ID    = "TEST_STRIPE_CUSTOMER_ACCOUNT_ID"
	TEST_STRIPE_KEY_FILE               = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/Stripe-TEST-Key.json"
	TEST_STRIPE_TOKEN                  = "This is a fake token"
	TEST_TLS_CABundle_FQN              = "/Volumes/development-share/.keys/savup/STAR_savup_com/CAbundle.crt"
	TEST_TLS_CERTIFICATE_FQN           = "/Volumes/development-share/.keys/savup/STAR_savup_com/STAR_savup_com.crt"
	TEST_TLS_PRIVATE_KEY_FQN           = "/Volumes/development-share/.keys/savup/savup.com.key"
	TEST_TOKEN_INVALID                 = "Im and invalid token"
	TEST_USER_REGISTER_COLLECTION_PATH = "user_register/wire/2023-02-07"
	TEST_UNREADABLE_FQN                = "/Users/syacko/workspace/styh-dev/src/albert/savup-server/testing/files/unreadable-file.txt"
	TEST_URL_INVALID                   = "not a url"
	TEST_URL_VALID                     = "https://google.com"
	TEST_USER_POOL_ID                  = "us-west-2_AOiL0wX83"
	TEST_UUID_INVALID                  = "This is not a UUID"
	TEST_UUID_VALID                    = "3766880e-5077-48eb-b304-ae6c636f7ac5" //
	//
	// Floats and Numbers
	TEST_FLOAT_10_01  = 10.01
	TEST_FLOAT_123_01 = 123.01
	TEST_NUMBER_44    = 44
	TEST_NUMBER_23    = 23
	//
	// User Data
	TEST_USER_AREA_CODE                    = "650"
	TEST_USER_BANK_ACCOUNT_ID_1            = "wmdBb98Ax4Ip19PJLwoGsqXZrgzW19FEbmvXZ"
	TEST_USER_BANK_ACCOUNT_ID_2            = "TEST_EjNJJxxxxxxxxxxxxxx6hRk5j3NwVrc97ppJQ_TEST"
	TEST_USER_BIRTHDAY                     = "01/01/2023"
	TEST_USER_CITY                         = "TEST_CITY"
	TEST_USER_EMAIL                        = "scott.yacko@sty-holdings.com"
	TEST_USER_FEDERAL_TAX_ID               = "999999999"
	TEST_USER_FIRST_NAME                   = "FIRST"
	TEST_USER_LAST_NAME                    = "LAST"
	TEST_USER_PHONE_NUMBER                 = "5551212"
	TEST_USER_REQUESTOR_ID_F               = "TEST_USER_REQUESTOR_ID_%v"
	TEST_USER_STATE                        = "CA"
	TEST_USER_STREET_ADDRESS               = "123 Main Street\nSuite ABC"
	TEST_USER_TRANSFER_AMOUNT              = 999.99
	TEST_USER_ZIP                          = "95391"
	TEST_USERNAME                          = "savup_test_do_not_delete"
	TEST_USERNAME_SAVUP_REQUESTOR_ID       = "7d2a5b84-5336-43e9-ace1-ff33fc900710"
	TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE = "savup_test_do_not_delete"
	TEST_VERSION                           = "TEST_VERSION"
	//
	// Plaid Data
	TEST_PLAID_KEY_FILE = "/Users/syacko/workspace/styh-dev/src/albert/keys/development/.keys/plaid-dev-creds.json"
	//
	// Email Data
	//
	// Email Templates
	TEST_EMAIL_TEMPLATE_VERIFACTION_ID = "d-1b95eb212cb7460f9ada4d0db09a2b3f"
)
