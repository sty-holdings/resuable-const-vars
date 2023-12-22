// Package constants
/*
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
package constants

//goland:noinspection ALL
const (
	LOG_MESSAGE_PREFIX = "albert/SavUpServer"
	//
	FN_ACCOUNTS                           = "accounts"
	FN_ACCOUNT_ID                         = "account_id"
	FN_ACCOUNTS_STRIPE_BANK_ACCOUNT_TOKEN = "accounts.stripe_bank_account_tokens."
	FN_ADDED_PHONE                        = "added_phone"
	FN_ADDED_SAVINGS_GOAL                 = "added_savings_goal"
	FN_ADDITIONAL_TRANSFER_INFO           = "additional_transfer_info"
	FN_AGG_TOTAL_LOAN_AMOUNT              = "agg_total_loan_amount"
	FN_AGG_TWI_SCORE                      = "agg_twi_score"
	FN_AMOUNT_IN_PENNIES                  = "amount_in_pennies"
	FN_AREA_CODE                          = "area_code"
	FN_AREA_CODE_CUSTOM                   = "custom:areaCode"
	FN_AUDIENCE                           = "AUDIENCE"
	FN_AVAILABLE_FUNDS                    = "available_funds"
	FN_AWS_ACCESS_KEY_ID                  = "AWS_ACCESS_KEY_ID"
	FN_AWS_INFO_FQN                       = "AWS_Info_FQN"
	FN_AWS_PROFILE                        = "AWS_PROFILE"
	FN_AWS_REGION                         = "AWS_REGION"
	FN_AWS_SECRET_ACCESS_KEY              = "AWS_SECRET_ACCESS_KEY"
	FN_BALANCE                            = "Balance"
	FN_BIRTHDATE                          = "birthdate"
	FN_BUNDLE_ID                          = "bundle_id"
	FN_BUNDLE_TITLE                       = "bundle_title"
	FN_CERT_KID                           = "kid"
	FN_CITY                               = "city"
	FN_COGNITO_USERNAME                   = "cognito:username"
	FN_CREATE_TIMESTAMP                   = "create_timestamp"
	FN_DATASTORE                          = "datastore"
	FN_DESCRIPTION                        = "description"
	FN_DOCUMENT_ID                        = "document_id"
	FN_DURATION                           = "Duration"
	FN_EMAIL                              = "email"
	FN_EMAIL_VERIFIED                     = "email_verified"
	FN_EXPIRY_TIMESTAMP                   = "expiry_timestamp"
	FN_FAMILY_NAME                        = "family_name"
	FN_FIRST_NAME                         = "first_name"
	FN_FEDERAL_TAX_ID                     = "federal_tax_id"
	FN_FEDERAL_TAX_ID_IS_SET              = "federal_tax_id_is_set"
	FN_GIVEN_NAME                         = "given_name"
	FN_INSTITUTIONS                       = "institutions"
	FN_INSTITUTION_NAME                   = "institution_name"
	FN_INSTITUTION_ACCOUNT                = "institution_account"
	FN_IS_BUNDLE_LOCKED                   = "is_bundle_locked"
	FN_ISSUER                             = "issuer"
	FN_JWT                                = "JWT"
	FN_LAST_NAME                          = "last_name"
	FN_LAST_REFRESHED                     = "last_refreshed"
	FN_LAST_UPDATE_TIMESTAMP              = "last_update_timestamp"
	FN_LINKED_BANK                        = "linked_bank"
	FN_LOAN_TYPE                          = "loan_type"
	FN_LOANED_AMOUNT_INVESTED             = "loaned_amount_invested"
	FN_LOANED_AMOUNT_RETURNED             = "loaned_amount_returned"
	FN_LOANS                              = "loans"
	FN_LOCKUP_END_DATE                    = "lockup_end_date"
	FN_LOCKUP_MONTHS                      = "lockup_months"
	FN_LOCKUP_START_DATE                  = "lockups_start_date"
	FN_LOG_DIRECTORY                      = "log_directory"
	FN_MAX_ALLOCATION                     = "max_allocation"
	FN_MIN_ALLOCATION                     = "min_allocation"
	FN_NAVIGATION                         = "navigation"
	FN_NICKNAME                           = "nickname"
	FN_OFFERED_INTEREST_RATE              = "offered_interest_rate"
	FN_OFFICIAL_NAME                      = "official_name"
	FN_PAYMENT_FREQENCY                   = "payment_frequency"
	FN_PERIOD                             = "peroid"
	FN_PLAID_ACCESS_TOKEN                 = "plaid_access_token"
	FN_PLAID_ACCOUNT                      = "plaid_account"
	FN_PLAID_ACCOUNTS                     = "plaid_accounts"
	FN_PLAID_INFO_FQN                     = "Plaid_Key_FQN"
	FN_PLAID_ITEM_ID                      = "plaid_item_id"
	FN_PHONE_NUMBER                       = "phone_number"
	FN_PHONE_NUMBER_CUSTOM                = "custom:phoneNumber"
	FN_PHONE_VERIFIED                     = "phone_verified"
	FN_PID_DIRECTORY                      = "pid_directory"
	FN_PRIVATE_KEY                        = "PrivateKey"
	FN_PROVIDED_AML_PHOTO                 = "provided_aml_photo"
	FN_PURPOSE                            = "purpose"
	FN_REPORT_BALANCE                     = "report_balance"
	FN_REPORT_BALANCE_SOURCE              = "report_balance_source"
	FN_REQUESTOR_ID                       = "requestor_id"
	FN_RISK_RATING                        = "risk_rating"
	FN_SAVUP_TAKE                         = "savup_take"
	FN_SET_BANKER_PREFERENCES             = "set_banker_preferences"
	FN_SHORT_URL                          = "short_URL"
	FN_STATE                              = "state"
	FN_STATUS                             = "status"
	FN_STREET_ADDRESS                     = "street_address"
	FN_STRIPE_ACCESS_TOKEN                = "stripe_access_token"
	FN_STRIPE_CUSTOMER_ACCOUNT_ID         = "StripeCustomerAccountId"
	FN_STRIPE_LOCK                        = "stripe_lock"
	FN_SUB                                = "sub"
	FN_TOKEN_PAYLOAD                      = "tokenPayload"
	FN_TOTAL_OFFERING_SIZE                = "total_offering_size"
	FN_TRANSFER_DIRECTION                 = "transfer_direction"
	FN_TRANSFERRED_FUNDS                  = "transferred_funds"
	FN_TRANSFER_METHOD                    = "transfer_method"
	FN_TRANSFER_STATUS                    = "transfer_status"
	FN_TWI_RATE                           = "twi_rate"
	FN_UPDATED_ADDRESS                    = "updated_address"
	FN_USER_BUNDLE_ALLOCATED_AMOUNT       = "user_bundle_allocated_amount"
	FN_USER_BUNDLE_ALLOCATION_DATE        = "user_bundle_allocation_date"
	FN_USER_BUNDLE_LOCKED                 = "user_bundle_locked"
	FN_USER_CONFIRMED_EMAIL               = "confirmed_email"
	FN_USER_CONFIRMED_PHONEL              = "confirmed_phone"
	FN_USERNAME                           = "username"
	FN_UUID                               = "uuid"
	FN_WEB_ASSETS_URL                     = "Web_Assets_URL"
	FN_ZELLE_REQUEST_METHOD               = "zelle_request_method"
	FN_ZIPCODE                            = "zip_code"

	//
	// Statuses
	STATUS_COMPLETED = "COMPLETED"
	STATUS_CONFIRMED = "CONFIRMED"
	STATUS_FAILED    = "FAILED"
	STATUS_LINKED    = "LINKED"
	STATUS_SUCCESS   = "SUCCESS"
	STATUS_PENDING   = "PENDING"
	STATUS_ACTIVE    = "ACTIVE"
	STATUS_INACTIVE  = "INACTIVE"
	//
	// Purposes
	PURPOSE_EMAIL_VERIFICATION    = "verify"
	PURPOSE_BANK_VERIFICATION     = "bank"
	PURPOSE_TRANSFER_VERIFICATION = "transfer"
	//
	// General Labels
	ATTEMPT_LIMIT_EXCEEDED    = "Attempt limit exceeded"
	AUTHENTICATOR_SERVICE     = "AuthenticatorService"
	BYPASS                    = "bypass"
	CERTIFICATE_FQN           = "CertificateFQN"
	CHECK                     = "Check"
	DEBUG_MESSAGE             = "DEBUG_MESSAGE:"
	DEBUG_MODE                = "DEBUG"
	EMAIL                     = "EMAIL"
	EMPTY                     = ""
	EMPTY_WORD                = "EMPTY"
	EQUALS                    = "=="
	FALSE                     = "false"
	FIREBASE_CREDENTIALS_FQN  = "FirebaseCredentialsFQN"
	FORWARD_SLASH             = "/"
	GCP_CREDENTIALS_FQN       = "GCPCredentialsFQN"
	GOOD                      = "good"
	INTO                      = "into"
	INVALID                   = "INVALID"
	IS_EMPTY                  = " is empty"
	IS_MISSING                = " is missing"
	IS_OK                     = " is ok."
	IS_UNREADABLE             = " is unreadable"
	LINE_SEPARATOR            = "----------------"
	LOCAL_HOST                = "localhost"
	MESSAGE_LABEL             = "message"
	MIME_TYPE_HTML            = "HTML"
	MIME_TYPE_PLAIN           = "PLAIN"
	MISSING                   = "MISSING"
	MISSING_KEY               = "MISSING KEY"
	MISSING_VALUE             = "MISSING VALUE"
	NA                        = "N.A."
	NATS_CREDENTIALS_FQN      = "NATSCredsFQN"
	NATS_NON_TLS_CONNECTION   = "NON-TLS"
	NATS_TLS_CONNECTION       = "TLS "
	NIL                       = "nil"
	NONE                      = "none"
	NOT_FOUND                 = "not found"
	PHONE                     = "PHONE"
	PID_FILENAME              = "/server.pid"
	PLAID_CLIENT_ID_NAME      = "PLAID-CLIENT-ID"
	PLAID_KEY_FQN             = "PlaidKeyFQN"
	PLAID_SECRET_NAME         = "PLAID-SECRET"
	PROCESS_MANUALLY          = "Process transfer manually."
	PROTECTED                 = "*********"
	OUTOF                     = "out of"
	RELEASE_MODE              = "RELEASE"
	SENDGRID_KEY_FQN          = "SendGridKeyFQN"
	STRIPE                    = "Stripe"
	STRIPE_KEY_FQN            = "StripeKeyFQN"
	TLS_CABUNDLE_FQN          = "TLS:TLSCABundle"
	TLS_CERTIFICATE_FQN       = "TLS:TLSCert"
	TLS_PRIVATE_KEY_FQN       = "TLS:TLSKey"
	TRANSFER_INSTITUTION_NAME = "Transfer Bank:"
	VALID                     = "VALID"
	TRUE                      = "true"
	WIRE                      = "Wire"
	ZELLE                     = "Zelle"
	//
	// Datastores & Collections
	DATASTORE_BUNDLES                   = "bundles"
	DATASTORE_SHORT_URL                 = "short_urls"
	DATASTORE_USERS                     = "users"
	DATASTORE_USER_BUNDLES              = "user_bundles"
	DATASTORE_USER_INSTITUTIONS         = "user_institutions"
	DATASTORE_USER_INSTITUTION_ACCOUNTS = "user_institution_accounts"
	DATASTORE_USER_REGISTER             = "user_register"
	DATASTORE_USER_TO_DOS               = "user_to_dos"
	COLLECTION_INSTITUTIONS             = "INSTITUTIONS"
	COLLECTION_USER_GOALS               = "GOALS"
	COLLECTION_USER_TO_DO_LIST          = "TODOLIST"
	//
	// Mine
	MIME_HTML_TEXT  = "MIME-version: 1.0;\nContent-Transfer-Encoding: 8bit;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	MIME_PLAIN_TEXT = "MIME-version: 1.0;\nContent-Transfer-Encoding: 8bit;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	//
	// Environment
	ENVIRONMENT_DEVELOPMENT = "DEVELOPMENT"
	ENVIRONMENT_LOCAL       = "LOCAL"
	ENVIRONMENT_PRODUCTION  = "PRODUCTION"
	//
	// Messages
	// Prefixes
	MESSAGE_PREFIX_SAVUPPROD  = "SAVUPPROD"
	MESSAGE_PREFIX_SAVUPDEV   = "SAVUPDEV"
	MESSAGE_PREFIX_SAVUPLOCAL = "SAVUPLOCAL"
	//
	// SavUp server Command subjects
	COMMAND_DEBUG         = ".command.debug"
	COMMAND_LIST_MESSAGES = ".command.listMessages"
	COMMAND_SHUTDOWN      = ".command.shutdown"
	SAVUP_HTTP_SUFFIX     = ".savup-http"   // this must match the the binary name
	SAVUP_SERVER_SUFFIX   = ".savup-server" // this must match the the binary name
	//
	// SavUp application subjects
	MSG_BUNDLE_ALLOCATION_ADJUSTMENT   = ".bundleAllocationAdjustment"
	MSG_CHECK_BUNDLE_EXISTS            = ".checkBundleExists"
	MSG_CREATE_BUNDLE                  = ".createBundle"
	MSG_CREATE_SHORT_URL               = ".createShortURL"
	MSG_CREATE_USER                    = ".createUser"
	MSG_CUSTOMER_TRANSFER              = ".customerTransfer"
	MSG_FORGOT_USERNAME                = ".forgotUsername"
	MSG_GET_ALL_BUNDLES                = ".getAllBundles"
	MSG_GET_BUNDLE                     = ".getBundle"
	MSG_GET_INSTITUTION_BALANCES       = ".getInstitutionBalances"
	MSG_GET_INSTITUTION_INFO           = ".getInstitutionInfo"
	MSG_GET_BACKEND_INFO               = ".getBackendInfo"
	MSG_GET_LOAN_NOTE                  = ".getLoanNote"
	MSG_GET_TODO_LIST                  = ".getToDoList"
	MSG_GET_USER_BUNDLES               = ".getUserBundles"
	MSG_GET_USER_PROFILE               = ".getUserProfile"
	MSG_GET_USER_REGISTER              = ".getUserRegister"
	MSG_LIST_INSTITUTIONS              = ".listInstitutions"
	MSG_LIST_INSTITUTION_ACCOUNT_NAMES = ".listInstitutionAccountNames"
	MSG_PLAID_GET_LINK_TOKEN           = ".plaidGetLinkToken"
	MSG_PULL_USER                      = ".pullUser"
	MSG_PUSH_LINK_AND_CREATE_CUSTOMER  = ".pushLinkAndCreateCustomer"
	MSG_RESEND_VERIFICATION_EMAIL      = ".resendVerifyEmail"
	MSG_RESET_USER_PASSWORD            = ".resetUserPassword"
	MSG_SET_FEDERAL_TAX_ID             = ".setFederalTaxId"
	MSG_USER_VERIFICATION              = ".userVerification"
	MSG_UPDATE_TO_DO                   = ".updateToDo"
	MSG_UPDATE_USER_FUNDS              = ".updateUserFunds"
	MSG_UPDATE_USER_PROFILE            = ".updateUserProfile"
	//
	// 	Log Messages
	ENDING_EXECUTION = " - Ending execution!"
	//
	// Colors
	COLOR_RESET  = "\033[0m"
	COLOR_RED    = "\033[31m"
	COLOR_GREEN  = "\033[32m"
	COLOR_YELLOW = "\033[33m"
	COLOR_BLUE   = "\033[34m"
	COLOR_PURPLE = "\033[35m"
	COLOR_CYAN   = "\033[36m"
	COLOR_WHITE  = "\033[37m"
	//
	//  formatting
	FOUR_SPACES = "    "
	//
	//  Certificate & Crypto
	CERT_ED25519             = "ED25519"
	CERT_RSA                 = "RSA"
	CERT_RS265               = "RS256"
	CERT_ECDSACURVE          = "ECDSACURVE"
	CERT_ECDSACURVE_P224     = "P224"
	CERT_ECDSACURVE_P256     = "P256"
	CERT_ECDSACURVE_P384     = "P384"
	CERT_ECDSACURVE_P521     = "P521"
	CERT_PRIVATE_KEY         = "RSA PRIVATE KEY"
	CERT_PUBLIC_KEY          = "PUBLIC KEY"
	CERTIFICATE              = "CERTIFICATE"
	CERT_ISSUER              = "STY HOLDINGS INC"
	CERT_SAVUP_AUDIENCE      = "savup-f3343"
	CERT_SAVUPDEV_AUDIENCE   = "savup-development"
	CERT_ID_TOEKN_ISSUER     = "https://securetoken.google.com/savup-f3343"
	CERT_DEV_ID_TOEKN_ISSUER = "https://securetoken.google.com/savup-development"
	//
	// Token Types
	TOKEN_TYPE_ID     = "ID"
	TOKEN_TYPE_ACCESS = "ACCESS"
	//
	// Time
	DAY   = "D"
	MONTH = "M"
	YEAR  = "Y"
	//
	//  Loan Methods
	TRANFER_STRIPE = "STRIPE"
	TRANFER_WIRE   = "WIRE"
	TRANFER_CHECK  = "CHECK"
	TRANFER_ZELLE  = "ZELLE"
	TRANSFER_IN    = "IN"
	TRANSFER_OUT   = "OUT"
	//
	// JWT Payload Fields
	PAYLOAD_SUBJECT_FN      = "SUBJECT"
	PAYLOAD_CLAIMS_FN       = "CLAIMS"
	PAYLOAD_AUDIENCE_FN     = "AUDIENCE"
	PAYLOAD_REQUESTOR_ID_FN = "REQUESTOR_ID"
	PAYLOAD_EXPIRES_FN      = "EXPIRES"
	PAYLOAD_ISSUER_FN       = "ISSUER"
	PAYLOAD_ISSUED_AT_FN    = "ISSUED_AT"
	//
	// HTTP
	//
	// HTTP Domains
	HTTP_DOMAIN_SAVUP     = "savup.com"
	HTTP_DOMAIN_LOCALHOST = "localhost"
	HTTP_DOMAIN_API_LOCAL = "api-local." + HTTP_DOMAIN_SAVUP
	HTTP_DOMAIN_API_DEV   = "api-dev." + HTTP_DOMAIN_SAVUP
	HTTP_DOMAIN_API_PROD  = "api-prod." + HTTP_DOMAIN_SAVUP
	//
	// HTTP Protocols
	HTTP_PROTOCOL_NON_SECURE = "http"
	HTTP_PROTOCOL_SECURE     = "https"
	//
	//  HTTP Ports
	HTTP_PORT_SECURE     = 8443
	HTTP_PORT_NON_SECURE = 8080
	//
	// HTTP Methods
	HTTP_POST = "POST"
	HTTP_GET  = "GET"
	//
	//  HTTP Status codes
	HTTP_STATUS_200 = 200 // Success
	HTTP_STATUS_400 = 400 // Client Error
	HTTP_STATUS_401 = 401 // Unauthorized to access page
	HTTP_STATUS_404 = 404 // Page not found
	HTTP_STATUS_500 = 500 // Server Error
	//
	// HTTP Endpoints
	ENDPOINT_FORGOT_USERNAME     = "forgotusername"
	ENDPOINT_PULL_USER           = "pulluser"
	ENDPOINT_RESET_USER_PASSWORD = "resetuserpassword"
	ENDPOINT_RESEND_VERIFY_EMAIL = "resendverifyemail"
	ENDPOINT_VERIFY_EMAIL        = "verifyemail"
	ENDPOINT_GET_HTTP_INFO       = "gethttpinfo"
	ENDPOINT_GET_BACKEND_INFO    = "getbackendinfo"
	//
	// Authenicators
	AUTH_FIREBASE = "FIREBASE"
	AUTH_COGNITO  = "COGNITO"
	AUTH_TEST     = "TEST"
	//
	// 	Server Names
	SAVUP_HTTP   = "savup-http"
	SAVUP_SERVER = "savup-server"
	//
	// Display Fields
	DISPLAY_BUNDLE_ID_F    = "Bundle Id: %v"
	DISPLAY_REQUESTOR_ID_F = "Requestor Id: %v"
	//
	// output Modes
	OUTPUT_DISPLAY = "display"
	OUTPUT_LOG     = "log"
	//
	// To DO Items
	TO_DO_PROFILE  = "Update your profile. (Required)"
	TO_DO_TAX_INFO = "Enter your SSN or EIN."
	TO_DO_TRANSFER = "Lend money to SavUp"
	//
	// Mobile App Navigation
	MOBILE_APP_PROFILE  = "/contactDetails"
	MOBILE_APP_TAX_INFO = "/taxInfo"
	MOBILE_APP_TRANSFER = "/moveMoney"
	//
	// Postgres
	POSTGRES_SSL_MODE_DISABLE  = "disable"
	POSTGRES_SSL_MODE_ALLOW    = "allow"
	POSTGRES_SSL_MODE_PREFER   = "prefer"
	POSTGRES_SSL_MODE_REQUIRED = "require"
	POSTGRES__CONN_STRING      = "dbname=%v user=%v password=%v host=%v port=%v connect_timeout=%v sslmode=%v pool_max_conns=%v"
)
