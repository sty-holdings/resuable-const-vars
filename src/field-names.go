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
)

//goland:noinspection ALL
const (
	PAYLOAD_SUBJECT_FN      = "SUBJECT"
	PAYLOAD_CLAIMS_FN       = "CLAIMS"
	PAYLOAD_AUDIENCE_FN     = "AUDIENCE"
	PAYLOAD_REQUESTOR_ID_FN = "REQUESTOR_ID"
	PAYLOAD_EXPIRES_FN      = "EXPIRES"
	PAYLOAD_ISSUER_FN       = "ISSUER"
	PAYLOAD_ISSUED_AT_FN    = "ISSUED_AT"
)
