package confuse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func allOfSchemas(schemas ...Schema) Schema {
	return Schema{
		AllOf: schemas,
	}
}

func singleSchemaFromPattern(pattern string) Schema {
	return allOfSchemas(Schema{
		Pattern: pattern,
	})
}

func TestService_ValidateTagsToSchema(t *testing.T) {
	testCases := []struct {
		name           string
		validateTag    string
		expectedSchema Schema
		required       bool
	}{
		{
			name:        "combination of tags",
			validateTag: "required,eq=foo,gt=1,gte=1.5,lt=3,lte=2,oneof=foo bar",
			required:    true,
			expectedSchema: allOfSchemas(
				Schema{
					Pattern: "^foo$",
				},
				Schema{
					Pattern: "^[0-9]+$",
					Minimum: 1,
				},
				Schema{
					Pattern:          "^[0-9]+$",
					Minimum:          1.5,
					ExclusiveMinimum: false,
				},
				Schema{
					Pattern: "^[0-9]+$",
					Maximum: 3,
				},
				Schema{
					Pattern:          "^[0-9]+$",
					Maximum:          2,
					ExclusiveMaximum: false,
				},
				Schema{
					Enum: []interface{}{
						"foo",
						"bar",
					},
				}),
		},
		{
			name:           "alpha",
			validateTag:    "alpha",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z]+$"),
		},
		{
			name:           "alphanum",
			validateTag:    "alphanum",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z0-9]+$"),
		},
		{
			name:           "alphanumunicode",
			validateTag:    "alphanumunicode",
			expectedSchema: singleSchemaFromPattern("^\\p{L}+$"),
		},
		{
			name:           "alphaunicode",
			validateTag:    "alphaunicode",
			expectedSchema: singleSchemaFromPattern("^\\p{L}+$"),
		},
		{
			name:           "ascii",
			validateTag:    "ascii",
			expectedSchema: singleSchemaFromPattern("^\\p{ASCII}+$"),
		},
		{
			name:        "boolean",
			validateTag: "boolean",
			expectedSchema: allOfSchemas(Schema{
				Type: "boolean",
			}),
		},
		{
			name:           "contains",
			validateTag:    "contains=foo",
			expectedSchema: singleSchemaFromPattern("foo"),
		},
		{
			name:           "containsany",
			validateTag:    "containsany=foo",
			expectedSchema: singleSchemaFromPattern("[foo]"),
		},
		{
			name:           "containsrune",
			validateTag:    "containsrune=foo",
			expectedSchema: singleSchemaFromPattern("[foo]"),
		},
		{
			name:           "endsnotwith",
			validateTag:    "endsnotwith=foo",
			expectedSchema: singleSchemaFromPattern("[^foo]$"),
		},
		{
			name:           "endswith",
			validateTag:    "endswith=foo",
			expectedSchema: singleSchemaFromPattern("foo$"),
		},
		{
			name:           "excludes",
			validateTag:    "excludes=foo",
			expectedSchema: singleSchemaFromPattern("^(?!foo).*$"),
		},
		{
			name:           "excludesall",
			validateTag:    "excludesall=foo",
			expectedSchema: singleSchemaFromPattern("^(?!foo).*$"),
		},
		{
			name:           "excludesrune",
			validateTag:    "excludesrune=foo",
			expectedSchema: singleSchemaFromPattern("^(?!foo).*$"),
		},
		{
			name:           "lowercase",
			validateTag:    "lowercase",
			expectedSchema: singleSchemaFromPattern("^[a-z]+$"),
		},
		{
			name:           "multibyte",
			validateTag:    "multibyte",
			expectedSchema: singleSchemaFromPattern("[^\x00-\x7F]"),
		},
		{
			name:        "number",
			validateTag: "number",
			expectedSchema: allOfSchemas(Schema{
				Type: "number",
			}),
		},
		{
			name:           "numeric",
			validateTag:    "numeric",
			expectedSchema: singleSchemaFromPattern("^[0-9]+$"),
		},
		{
			name:           "printascii",
			validateTag:    "printascii",
			expectedSchema: singleSchemaFromPattern("^[\x20-\x7E]+$"),
		},
		{
			name:           "startsnotwith",
			validateTag:    "startsnotwith=foo",
			expectedSchema: singleSchemaFromPattern("^[^foo]"),
		},
		{
			name:           "startswith",
			validateTag:    "startswith=foo",
			expectedSchema: singleSchemaFromPattern("^foo"),
		},
		{
			name:           "uppercase",
			validateTag:    "uppercase",
			expectedSchema: singleSchemaFromPattern("^[A-Z]+$"),
		},
		{
			name:           "base64",
			validateTag:    "base64",
			expectedSchema: singleSchemaFromPattern("^[A-Za-z0-9+/]*={0,2}$"),
		},
		{
			name:           "base64url",
			validateTag:    "base64url",
			expectedSchema: singleSchemaFromPattern("^[A-Za-z0-9+/]*={0,2}$"),
		},
		{
			name:           "base64rawurl",
			validateTag:    "base64rawurl",
			expectedSchema: singleSchemaFromPattern("^[A-Za-z0-9-_]*$"),
		},
		{
			name:           "bic",
			validateTag:    "bic",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$"),
		},
		{
			name:           "bcp47_language_tag",
			validateTag:    "bcp47_language_tag",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*$"),
		},
		{
			name:           "btc_addr",
			validateTag:    "btc_addr",
			expectedSchema: singleSchemaFromPattern("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$"),
		},
		{
			name:           "btc_addr_bech32",
			validateTag:    "btc_addr_bech32",
			expectedSchema: singleSchemaFromPattern("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$"),
		},
		{
			name:           "credit_card",
			validateTag:    "credit_card",
			expectedSchema: singleSchemaFromPattern("^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9]{2})[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35[0-9]{3})[0-9]{11})$"),
		},
		{
			name:           "mongodb",
			validateTag:    "mongodb",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z0-9-_]+$"),
		},
		{
			name:           "spicedb",
			validateTag:    "spicedb",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z0-9-_]+$"),
		},
		{
			name:           "cron",
			validateTag:    "cron",
			expectedSchema: singleSchemaFromPattern("^((\\*(/[0-9]+)?|[0-9,-]+(,[0-9,-]+)*) ?){5}$"),
		},
		{
			name:        "datetime",
			validateTag: "datetime",
			expectedSchema: allOfSchemas(Schema{
				Format: "date-time",
			}),
		},
		{
			name:           "e164",
			validateTag:    "e164",
			expectedSchema: singleSchemaFromPattern("^\\+[1-9]\\d{1,14}$"),
		},
		{
			name:        "email",
			validateTag: "email",
			expectedSchema: allOfSchemas(Schema{
				Format: "email",
			}),
		},
		{
			name:           "eth_addr",
			validateTag:    "eth_addr",
			expectedSchema: singleSchemaFromPattern("^0x[a-fA-F0-9]{40}$"),
		},
		{
			name:           "hexadecimal",
			validateTag:    "hexadecimal",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]+$"),
		},
		{
			name:           "hexcolor",
			validateTag:    "hexcolor",
			expectedSchema: singleSchemaFromPattern("^#?([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$"),
		},
		{
			name:           "hsl",
			validateTag:    "hsl",
			expectedSchema: singleSchemaFromPattern("^hsl\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*\\)$"),
		},
		{
			name:           "hsla",
			validateTag:    "hsla",
			expectedSchema: singleSchemaFromPattern("^hsla\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*,\\s*(0?.\\d+|1)\\s*\\)$"),
		},
		{
			name:           "html",
			validateTag:    "html",
			expectedSchema: singleSchemaFromPattern("^<([a-z]+)([^<]+)*(?:>(.*)<\\/\\1>|\\s+\\/>)$"),
		},
		{
			name:           "html_encoded",
			validateTag:    "html_encoded",
			expectedSchema: singleSchemaFromPattern("&([a-z]+|#[0-9]{1,6}|#x[0-9a-fA-F]{1,6});"),
		},
		{
			name:           "isbn",
			validateTag:    "isbn",
			expectedSchema: singleSchemaFromPattern("^(?:ISBN(?:-1[03])?:? )?(?=[-0-9 ]{17}$|[-0-9X ]{13}$|[0-9X]{10}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9X]$"),
		},
		{
			name:           "isbn10",
			validateTag:    "isbn10",
			expectedSchema: singleSchemaFromPattern("^(?:ISBN(?:-10)?:? )?(?=[-0-9X ]{13}$|[0-9X]{10}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9X]$"),
		},
		{
			name:           "isbn13",
			validateTag:    "isbn13",
			expectedSchema: singleSchemaFromPattern("^(?:ISBN(?:-13)?:? )?(?=[-0-9 ]{17}$|[0-9]{13}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9]$"),
		},
		{
			name:           "issn",
			validateTag:    "issn",
			expectedSchema: singleSchemaFromPattern("^\\d{4}-?\\d{3}[\\dX]$"),
		},
		{
			name:           "iso3166_1_alpha2",
			validateTag:    "iso3166_1_alpha2",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{2}$"),
		},
		{
			name:           "iso3166_1_alpha3",
			validateTag:    "iso3166_1_alpha3",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{3}$"),
		},
		{
			name:           "iso3166_1_numeric",
			validateTag:    "iso3166_1_numeric",
			expectedSchema: singleSchemaFromPattern("^[0-9]{3}$"),
		},
		{
			name:           "iso3166_2",
			validateTag:    "iso3166_2",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{2}-[A-Z0-9]{1,3}$"),
		},
		{
			name:           "iso4217",
			validateTag:    "iso4217",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{3}$"),
		},
		{
			name:           "jwt",
			validateTag:    "jwt",
			expectedSchema: singleSchemaFromPattern("^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.([A-Za-z0-9-_]+)?$"),
		},
		{
			name:           "latitude",
			validateTag:    "latitude",
			expectedSchema: singleSchemaFromPattern("^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"),
		},
		{
			name:           "longitude",
			validateTag:    "longitude",
			expectedSchema: singleSchemaFromPattern("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"),
		},
		{
			name:           "luhn_checksum",
			validateTag:    "luhn_checksum",
			expectedSchema: singleSchemaFromPattern("^[0-9]{12,19}$"),
		},
		{
			name:           "postcode_iso3166_alpha2",
			validateTag:    "postcode_iso3166_alpha2",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{2}-[0-9]{5}$"),
		},
		{
			name:           "postcode_iso3166_alpha2_field",
			validateTag:    "postcode_iso3166_alpha2_field",
			expectedSchema: singleSchemaFromPattern("^[A-Z]{2}$"),
		},
		{
			name:           "rgb",
			validateTag:    "rgb",
			expectedSchema: singleSchemaFromPattern("^rgb\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*\\)$"),
		},
		{
			name:           "rgba",
			validateTag:    "rgba",
			expectedSchema: singleSchemaFromPattern("^rgba\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*(0?.\\d+|1)\\s*\\)$"),
		},
		{
			name:           "ssn",
			validateTag:    "ssn",
			expectedSchema: singleSchemaFromPattern("^(?!000|666)[0-8]\\d{2}-(?!00)\\d{2}-(?!0000)\\d{4}$"),
		},
		{
			name:           "timezone",
			validateTag:    "timezone",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z_]+/[a-zA-Z_]+$"),
		},
		{
			name:        "uuid",
			validateTag: "uuid",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid3",
			validateTag: "uuid3",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid3_rfc4122",
			validateTag: "uuid3_rfc4122",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid4",
			validateTag: "uuid4",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid4_rfc4122",
			validateTag: "uuid4_rfc4122",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid5",
			validateTag: "uuid5",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid5_rfc4122",
			validateTag: "uuid5_rfc4122",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:        "uuid_rfc4122",
			validateTag: "uuid_rfc4122",
			expectedSchema: allOfSchemas(Schema{
				Format: "uuid",
			}),
		},
		{
			name:           "md4",
			validateTag:    "md4",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{32}$"),
		},
		{
			name:           "md5",
			validateTag:    "md5",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{32}$"),
		},
		{
			name:           "sha256",
			validateTag:    "sha256",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{64}$"),
		},
		{
			name:           "sha384",
			validateTag:    "sha384",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{96}$"),
		},
		{
			name:           "sha512",
			validateTag:    "sha512",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{128}$"),
		},
		{
			name:           "ripemd128",
			validateTag:    "ripemd128",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{32}$"),
		},
		{
			name:           "ripemd160",
			validateTag:    "ripemd160",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{32}$"),
		},
		{
			name:           "tiger128",
			validateTag:    "tiger128",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{32}$"),
		},
		{
			name:           "tiger160",
			validateTag:    "tiger160",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{40}$"),
		},
		{
			name:           "tiger192",
			validateTag:    "tiger192",
			expectedSchema: singleSchemaFromPattern("^[a-fA-F0-9]{40}$"),
		},
		{
			name:           "semver",
			validateTag:    "semver",
			expectedSchema: singleSchemaFromPattern("^v?(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"),
		},
		{
			name:           "ulid",
			validateTag:    "ulid",
			expectedSchema: singleSchemaFromPattern("^[0-9A-Z]{26}$"),
		},
		{
			name:           "cve",
			validateTag:    "cve",
			expectedSchema: singleSchemaFromPattern("^CVE-[0-9]{4}-[0-9]{4,}$"),
		},
		{
			name:           "cidr",
			validateTag:    "cidr",
			expectedSchema: singleSchemaFromPattern("^([0-9]{1,3}\\.){3}[0-9]{1,3}\\/[0-9]{1,2}$"),
		},
		{
			name:           "cidrv4",
			validateTag:    "cidrv4",
			expectedSchema: singleSchemaFromPattern("^([0-9]{1,3}\\.){3}[0-9]{1,3}\\/[0-9]{1,2}$"),
		},
		{
			name:           "cidrv6",
			validateTag:    "cidrv6",
			expectedSchema: singleSchemaFromPattern("^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\/[0-9]{1,3}$"),
		},
		{
			name:           "datauri",
			validateTag:    "datauri",
			expectedSchema: singleSchemaFromPattern("^data:[a-z]+\\/[a-z0-9-+.]+(;[a-z-]+=[a-z0-9-]+)*;base64,[a-z0-9!$&',()*+,;=\\-._~:@/?%\\s]*$"),
		},
		{
			name:        "fqdn",
			validateTag: "fqdn",
			expectedSchema: allOfSchemas(Schema{
				Format: "hostname",
			}),
		},
		{
			name:        "hostname",
			validateTag: "hostname",
			expectedSchema: allOfSchemas(Schema{
				Format: "hostname",
			}),
		},
		{
			name:        "hostname_port",
			validateTag: "hostname_port",
			expectedSchema: allOfSchemas(Schema{
				Format: "hostname",
			}),
		},
		{
			name:        "hostname_rfc1123",
			validateTag: "hostname_rfc1123",
			expectedSchema: allOfSchemas(Schema{
				Format: "hostname",
			}),
		},
		{
			name:        "ip",
			validateTag: "ip",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv4",
			}),
		},
		{
			name:        "ip_addr",
			validateTag: "ip_addr",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv4",
			}),
		},
		{
			name:        "ipv4",
			validateTag: "ipv4",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv4",
			}),
		},
		{
			name:        "ip4_addr",
			validateTag: "ip4_addr",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv4",
			}),
		},
		{
			name:        "ipv6",
			validateTag: "ipv6",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv6",
			}),
		},
		{
			name:        "ip6_addr",
			validateTag: "ip6_addr",
			expectedSchema: allOfSchemas(Schema{
				Format: "ipv6",
			}),
		},
		{
			name:           "mac",
			validateTag:    "mac",
			expectedSchema: singleSchemaFromPattern("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"),
		},
		{
			name:           "tcp4_addr",
			validateTag:    "tcp4_addr",
			expectedSchema: singleSchemaFromPattern("^([0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]+$"),
		},
		{
			name:           "tcp6_addr",
			validateTag:    "tcp6_addr",
			expectedSchema: singleSchemaFromPattern("^\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]:[0-9]+$"),
		},
		{
			name:           "tcp_addr",
			validateTag:    "tcp_addr",
			expectedSchema: singleSchemaFromPattern("^(([0-9]{1,3}\\.){3}[0-9]{1,3}|\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]):[0-9]+$"),
		},
		{
			name:           "udp4_addr",
			validateTag:    "udp4_addr",
			expectedSchema: singleSchemaFromPattern("^([0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]+$"),
		},
		{
			name:           "udp6_addr",
			validateTag:    "udp6_addr",
			expectedSchema: singleSchemaFromPattern("^\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]:[0-9]+$"),
		},
		{
			name:           "udp_addr",
			validateTag:    "udp_addr",
			expectedSchema: singleSchemaFromPattern("^(([0-9]{1,3}\\.){3}[0-9]{1,3}|\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]):[0-9]+$"),
		},
		{
			name:           "unix_addr",
			validateTag:    "unix_addr",
			expectedSchema: singleSchemaFromPattern("^/.+/.+"),
		},
		{
			name:        "uri",
			validateTag: "uri",
			expectedSchema: allOfSchemas(Schema{
				Format: "uri",
			}),
		},
		{
			name:        "url",
			validateTag: "url",
			expectedSchema: allOfSchemas(Schema{
				Format: "uri",
			}),
		},
		{
			name:        "http_url",
			validateTag: "http_url",
			expectedSchema: allOfSchemas(Schema{
				Format: "uri",
			}),
		},
		{
			name:           "url_encoded",
			validateTag:    "url_encoded",
			expectedSchema: singleSchemaFromPattern("^[a-zA-Z0-9-_.!~*'()]+$"),
		},
		{
			name:           "urn_rfc2141",
			validateTag:    "urn_rfc2141",
			expectedSchema: singleSchemaFromPattern("^urn:[a-zA-Z0-9][a-zA-Z0-9-]{0,31}:([a-zA-Z0-9()+,\\-.:=@;$_!*'%/?#]|%[0-9a-fA-F]{2})+$"),
		},
		{
			name:        "required",
			validateTag: "required",
			required:    true,
		},
		{
			name:           "eq",
			validateTag:    "eq=foo",
			expectedSchema: singleSchemaFromPattern("^foo$"),
		},
		{
			name:           "eq_ignore_case",
			validateTag:    "eq_ignore_case=foo",
			expectedSchema: singleSchemaFromPattern("(?i)^foo$"),
		},
		{
			name:        "gt",
			validateTag: "gt=1",
			expectedSchema: allOfSchemas(Schema{
				Pattern: "^[0-9]+$",
				Minimum: 1,
			}),
		},
		{
			name:        "gte",
			validateTag: "gte=1",
			expectedSchema: allOfSchemas(Schema{
				Pattern:          "^[0-9]+$",
				Minimum:          1,
				ExclusiveMinimum: false,
			}),
		},
		{
			name:        "lt",
			validateTag: "lt=1",
			expectedSchema: allOfSchemas(Schema{
				Pattern: "^[0-9]+$",
				Maximum: 1,
			}),
		},
		{
			name:        "lte",
			validateTag: "lte=1",
			expectedSchema: allOfSchemas(Schema{
				Pattern:          "^[0-9]+$",
				Maximum:          1,
				ExclusiveMaximum: false,
			}),
		},
		{
			name:           "ne",
			validateTag:    "ne=foo",
			expectedSchema: singleSchemaFromPattern("^(?!foo).*$"),
		},
		{
			name:           "ne_ignore_case",
			validateTag:    "ne_ignore_case=foo",
			expectedSchema: singleSchemaFromPattern("(?i)^(?!foo).*$"),
		},
		{
			name:        "len",
			validateTag: "len=1",
			expectedSchema: allOfSchemas(Schema{
				Minimum: 1,
				Maximum: 1,
			}),
		},
		{
			name:        "max",
			validateTag: "max=1",
			expectedSchema: allOfSchemas(Schema{
				Maximum:          1,
				ExclusiveMaximum: true,
			}),
		},
		{
			name:        "min",
			validateTag: "min=1",
			expectedSchema: allOfSchemas(Schema{
				Minimum:          1,
				ExclusiveMinimum: true,
			}),
		},
		{
			name:        "oneof",
			validateTag: "oneof=foo|bar",
			expectedSchema: allOfSchemas(Schema{
				Enum: []interface{}{"foo", "bar"},
			}),
		},
		{
			name:        "iscolor",
			validateTag: "iscolor",
			expectedSchema: allOfSchemas(Schema{
				OneOf: []Schema{
					singleSchemaFromPattern("^#?([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$"),
					singleSchemaFromPattern("^hsl\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*\\)$"),
					singleSchemaFromPattern("^hsla\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*,\\s*(0?.\\d+|1)\\s*\\)$"),
					singleSchemaFromPattern("^rgb\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*\\)$"),
					singleSchemaFromPattern("^rgba\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*(0?.\\d+|1)\\s*\\)$"),
				},
			}),
		},
		{
			name:        "country_code",
			validateTag: "country_code",
			expectedSchema: allOfSchemas(Schema{
				OneOf: []Schema{
					singleSchemaFromPattern("^[A-Z]{2}$"),
					singleSchemaFromPattern("^[A-Z]{3}$"),
					singleSchemaFromPattern("^[0-9]{3}$"),
				},
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := &Service{
				ShouldValidate: true,
			}
			schema, required := s.validateTagsToSchema(tc.validateTag)
			require.Equal(t, tc.expectedSchema, schema)
			require.Equal(t, tc.required, required)
		})
	}
}
