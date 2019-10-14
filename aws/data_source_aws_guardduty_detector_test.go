package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAWSGuarddutyDetectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsGuarddutyDetectorBasicResourceConfig(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccAwsGuarddutyDetectorBasicResourceDataConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "id", "aws_guardduty_detector.test", "id"),
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "status", "aws_guardduty_detector.test", "status"),
					resource.TestMatchResourceAttr("data.aws_guardduty_detector.test", "service_role_arn", regexp.MustCompile("^arn:aws:iam::[0-9]{12}:role/aws-service-role/guardduty.amazonaws.com/AWSServiceRoleForAmazonGuardDuty")),
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "finding_publishing_frequency", "aws_guardduty_detector.test", "finding_publishing_frequency"),
				),
			},
		},
	})
}

func TestAccAWSGuarddutyDetectorDataSource_explicit(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAwsGuarddutyDetectorExplicitConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "id", "aws_guardduty_detector.test", "id"),
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "status", "aws_guardduty_detector.test", "status"),
					resource.TestMatchResourceAttr("data.aws_guardduty_detector.test", "service_role_arn", regexp.MustCompile("^arn:aws:iam::[0-9]{12}:role/aws-service-role/guardduty.amazonaws.com/AWSServiceRoleForAmazonGuardDuty")),
					resource.TestCheckResourceAttrPair("data.aws_guardduty_detector.test", "finding_publishing_frequency", "aws_guardduty_detector.test", "finding_publishing_frequency"),
				),
			},
		},
	})
}

func testAccAwsGuarddutyDetectorBasicResourceConfig() string {
	return fmt.Sprintf(`
resource "aws_guardduty_detector" "test" {}
`)
}

func testAccAwsGuarddutyDetectorBasicResourceDataConfig() string {
	return fmt.Sprintf(`
resource "aws_guardduty_detector" "test" {}

data "aws_guardduty_detector" "test" {}
`)
}

func testAccAwsGuarddutyDetectorExplicitConfig() string {
	return fmt.Sprintf(`
resource "aws_guardduty_detector" "test" {}

data "aws_guardduty_detector" "test" {
	id = "${aws_guardduty_detector.test.id}"
}
`)
}
