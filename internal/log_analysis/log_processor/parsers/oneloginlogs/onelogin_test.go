package oneloginlogs

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/require"

	"github.com/panther-labs/panther/internal/log_analysis/log_processor/parsers/testutil"
	"github.com/panther-labs/panther/internal/log_analysis/log_processor/parsers/timestamp"
)

func TestOneLoginAuthEvent(t *testing.T) {
	// nolint: lll
	log := `{"create":{"_id":"6ce16f26-b25d-4070-9a30-4f259675a745"},"event_timestamp":"2020-04-15 22:51:24 UTC","otp_device_id":1159999,"user_id":50361999,"event_type_id":1400,"actor_user_name":"Jon Bixton","notes":"Authentication method: OTP (Valid OTP). Factor name: Duo Security","actor_system":"","uuid":"6ce16f26-b25d-4070-9a30-4f259675a745","user_name":"Jon Bixton","ipaddr":"136.25.254.254","account_id":142999,"authentication_factor_type":12,"authentication_factor_description":"Duo Security","user_agent":"Swagger-Codegen/1.0.0/ruby","authentication_factor_id":20402,"actor_user_id":50361999,"otp_device_name":"Duo Duo Security"}`
	expectedTime := time.Date(2020, 4, 15, 22, 51, 24, 0, time.UTC)
	expectedEvent := &OneLogin{
		EventTimestamp:                  (*timestamp.OneLoginTimestamp)(&expectedTime),
		OtpDeviceID:                     aws.Int(1159999),
		UserID:                          aws.Int(50361999),
		EventTypeID:                     aws.Int(1400),
		ActorSystem:                     aws.String(""),
		ActorUserName:                   aws.String("Jon Bixton"),
		Notes:                           aws.String("Authentication method: OTP (Valid OTP). Factor name: Duo Security"),
		UUID:                            aws.String("6ce16f26-b25d-4070-9a30-4f259675a745"),
		UserName:                        aws.String("Jon Bixton"),
		IPAddr:                          aws.String("136.25.254.254"),
		AccountID:                       aws.Int(142999),
		AuthenticationFactorType:        aws.Int(12),
		AuthenticationFactorDescription: aws.String("Duo Security"),
		UserAgent:                       aws.String("Swagger-Codegen/1.0.0/ruby"),
		AuthenticationFactorID:          aws.Int(20402),
		ActorUserID:                     aws.Int(50361999),
		OtpDeviceName:                   aws.String("Duo Duo Security"),
	}
	expectedEvent.SetCoreFields("OneLogin.Events", (*timestamp.RFC3339)(&expectedTime), expectedEvent)
	expectedEvent.AppendAnyIPAddress("136.25.254.254")
	parser := (&OneLoginParser{}).New()
	logs, err := parser.Parse(log)
	testutil.EqualPantherLog(t, expectedEvent.Log(), logs, err)
}

func TestOneLoginPrivilegeEvent(t *testing.T) {
	// nolint: lll
	log := `{"create":{"_id":"aca33051-d938-40cf-8dec-c382e397b4d7"},"event_timestamp":"2020-04-14 23:34:15 UTC","ipaddr":"136.25.254.254","privilege_id":338999,"privilege_name":"Manage Panther Cognito (Test)","account_id":142999,"event_type_id":72,"uuid":"aca33051-d938-40cf-8dec-c382e397b4d7","app_name":"Panther Cognito (Test)","actor_user_name":"Jon Bixton","user_id":50361999,"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36","user_name":"Jon Bixton","app_id":1120551,"actor_user_id":50361999,"actor_system":""}`
	expectedTime := time.Date(2020, 4, 14, 23, 34, 15, 0, time.UTC)
	expectedEvent := &OneLogin{
		EventTimestamp: (*timestamp.OneLoginTimestamp)(&expectedTime),
		UserID:         aws.Int(50361999),
		UserName:       aws.String("Jon Bixton"),
		EventTypeID:    aws.Int(72),
		ActorSystem:    aws.String(""),
		ActorUserName:  aws.String("Jon Bixton"),
		UUID:           aws.String("aca33051-d938-40cf-8dec-c382e397b4d7"),
		IPAddr:         aws.String("136.25.254.254"),
		AccountID:      aws.Int(142999),
		// nolint: lll
		UserAgent:     aws.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		ActorUserID:   aws.Int(50361999),
		AppID:         aws.Int(1120551),
		AppName:       aws.String("Panther Cognito (Test)"),
		PrivilegeID:   aws.Int(338999),
		PrivilegeName: aws.String("Manage Panther Cognito (Test)"),
	}
	expectedEvent.SetCoreFields("OneLogin.Events", (*timestamp.RFC3339)(&expectedTime), expectedEvent)
	expectedEvent.AppendAnyIPAddress("136.25.254.254")
	parser := (&OneLoginParser{}).New()
	logs, err := parser.Parse(log)
	testutil.EqualPantherLog(t, expectedEvent.Log(), logs, err)
}

func TestOneLoginDirectorySync(t *testing.T) {
	// nolint: lll
	log := `{"create":{"_id":"5c103afb-bcca-4d5b-8eae-a23e675cae82"},"event_type_id":14,"uuid":"5c103afb-bcca-4d5b-8eae-a23e67588888","resolution": 10,"event_timestamp":"2020-04-09 18:37:33 UTC","directory_sync_run_id":1586458888,"account_id":142999,"actor_user_name":"G Suite","actor_system":"G Suite","ipaddr":"","user_agent":"","user_id":53642999,"user_name":"Jon Bixton","notes":"Updated user"}`
	expectedTime := time.Date(2020, 4, 9, 18, 37, 33, 0, time.UTC)
	expectedEvent := &OneLogin{
		Resolution:         aws.Int(10),
		EventTimestamp:     (*timestamp.OneLoginTimestamp)(&expectedTime),
		DirectorySyncRunID: aws.Int(1586458888),
		UserID:             aws.Int(53642999),
		UserName:           aws.String("Jon Bixton"),
		EventTypeID:        aws.Int(14),
		ActorSystem:        aws.String("G Suite"),
		ActorUserName:      aws.String("G Suite"),
		UUID:               aws.String("5c103afb-bcca-4d5b-8eae-a23e67588888"),
		IPAddr:             aws.String(""),
		AccountID:          aws.Int(142999),
		UserAgent:          aws.String(""),
		Notes:              aws.String("Updated user"),
	}
	expectedEvent.SetCoreFields("OneLogin.Events", (*timestamp.RFC3339)(&expectedTime), expectedEvent)
	parser := (&OneLoginParser{}).New()
	logs, err := parser.Parse(log)
	testutil.EqualPantherLog(t, expectedEvent.Log(), logs, err)
}

func TestOneLoginRiskReasons(t *testing.T) {
	// keep sure we parse logs with risk_reason set correctly
	logs := []string{
		`{"create":{"_id":"9d038ce7-12a9-4548-9835-6c07d13dcb63"},"resource_type_id":null,"resolved_at":null,"object_id":null,"job_name":null,"authentication_factor_type":null,"radius_config_id":null,"mapping_id":null,"client_id":null,"risk_cookie_id":"dccab460dde68a296872a01204976b5e6a755671aa6234d14ccedbea10fd4173","param":null,"proxy_ip":null,"authentication_factor_description":null,"directory_id":null,"risk_reasons":"Boulder, Colorado, United States is a new location (20%)\nAccessed from a new IP address (20%)\nUnexpectedly high velocity (4587.45 km/hr) (20%)\nInfrequent access from 70.39.110.83 (20%)\nInfrequent access from Boulder, Colorado, United States (20%)","api_credential_name":null,"safe_to_unescape":null,"report_id":null,"app_name":null,"proxy_agent_name":null,"user_field_name":null,"otp_device_id":null,"directory_name":null,"notes":"Authentication method: Login App.\nTransitions: desktop_login -> username -> password -> mfa_login\nCorrelation_id: fc7f0a67-6970-4919-8a73-e5a3b4c76189","custom_message":null,"policy_name":null,"note_id":null,"user_field_id":null,"user_id":73503145,"service_job_name":null,"assumed_by_superadmin_or_reseller":null,"assuming_acting_user_id":null,"otp_device_name":null,"authentication_factor_id":null,"browser_fingerprint":"0335332a287a43e05e0e7cd8ec9eb44c","adc_name":null,"user_name":"Somebody","trusted_idp_id":null,"trusted_idp_name":null,"service_directory_id":null,"actor_user_id":73503145,"note_title":null,"login_name":null,"adc_id":null,"certificate_name":null,"privilege_id":null,"solved":null,"uuid":"9d038ce7-12a9-4548-9835-6c07d13dcb63","directory_sync_run_id":null,"proxy_agent_id":null,"service_job_id":null,"app_id":null,"error_description":null,"actor_system":"","privilege_name":null,"imported_user_id":null,"job_id":null,"role_id":null,"policy_type":null,"actor_user_name":"Somebody","account_id":148935,"resolution":null,"event_type_id":5,"risk_score":30,"report_name":null,"resolved_by_user_id":null,"imported_user_name":null,"group_name":null,"ipaddr":"70.39.110.83","group_id":null,"policy_id":null,"entity":null,"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36","certificate_id":null,"mapping_name":null,"task_name":null,"login_id":null,"role_name":null,"event_timestamp":"2020-08-21 00:32:38 UTC","radius_config_name":null}`, // nolint: lll
		`{"create":{"_id":"5cb13a5c-2280-4155-8b70-45a2a8c40e4d"},"resource_type_id":null,"resolved_at":null,"object_id":null,"job_name":null,"authentication_factor_type":null,"radius_config_id":null,"mapping_id":null,"client_id":null,"risk_cookie_id":"dccab460dde68a296872a01204976b5e6a755671aa6234d14ccedbea10fd4173","param":null,"proxy_ip":null,"authentication_factor_description":null,"directory_id":null,"risk_reasons":"Boulder, Colorado, United States is a new location (20%)\nAccessed from a new IP address (20%)\nUnexpectedly high velocity (4587.45 km/hr) (20%)\nInfrequent access from 70.39.110.83 (20%)\nInfrequent access from Boulder, Colorado, United States (20%)","api_credential_name":null,"safe_to_unescape":null,"report_id":null,"app_name":null,"proxy_agent_name":null,"user_field_name":null,"otp_device_id":null,"directory_name":null,"notes":null,"custom_message":null,"policy_name":null,"note_id":null,"user_field_id":null,"user_id":73503145,"service_job_name":null,"assumed_by_superadmin_or_reseller":null,"assuming_acting_user_id":null,"otp_device_name":null,"authentication_factor_id":null,"browser_fingerprint":"0335332a287a43e05e0e7cd8ec9eb44c","adc_name":null,"user_name":"Somebody","trusted_idp_id":null,"trusted_idp_name":null,"service_directory_id":null,"actor_user_id":73503145,"note_title":null,"login_name":null,"adc_id":null,"certificate_name":null,"privilege_id":null,"solved":null,"uuid":"5cb13a5c-2280-4155-8b70-45a2a8c40e4d","directory_sync_run_id":null,"proxy_agent_id":null,"service_job_id":null,"app_id":null,"error_description":null,"actor_system":"","privilege_name":null,"imported_user_id":null,"job_id":null,"role_id":null,"policy_type":null,"actor_user_name":"Somebody","account_id":148935,"resolution":null,"event_type_id":1001,"risk_score":30,"report_name":null,"resolved_by_user_id":null,"imported_user_name":null,"group_name":null,"ipaddr":"70.39.110.83","group_id":null,"policy_id":null,"entity":null,"user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36","certificate_id":null,"mapping_name":null,"task_name":null,"login_id":null,"role_name":null,"event_timestamp":"2020-08-21 00:32:29 UTC","radius_config_name":null}`,                                                                                                                                                    // nolint: lll
	}
	parser := (&OneLoginParser{}).New()
	for _, log := range logs {
		parsedLogs, err := parser.Parse(log)
		require.NoError(t, err, log)
		require.Len(t, parsedLogs, 1)                                         // parsing worked
		require.NotNil(t, parsedLogs[0].Event().(*OneLogin).RiskReasons, log) // risk_reason populated
	}
}
