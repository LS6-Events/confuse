package confuse

import (
	"strconv"
	"strings"
)

func (s *Service) validateTagsToSchema(validateTag string) (topLevel Schema, required bool) {
	if s.ShouldValidate {
		validation := strings.Split(validateTag, ",")
		for _, v := range validation {
			var levelRequired bool
			split := strings.Split(v, "=")
			key := split[0]

			var value string
			if len(split) > 1 {
				value = split[1]
			}

			var schema Schema
			switch key {
			// Strings
			case "alpha":
				schema.Pattern = "^[a-zA-Z]+$"
			case "alphanum":
				schema.Pattern = "^[a-zA-Z0-9]+$"
			case "alphanumunicode", "alphaunicode":
				schema.Pattern = "^\\p{L}+$"
			case "ascii":
				schema.Pattern = "^\\p{ASCII}+$"
			case "boolean":
				schema.Type = "boolean"
			case "contains":
				schema.Pattern = value
			case "containsany", "containsrune":
				schema.Pattern = "[" + value + "]"
			case "endsnotwith":
				schema.Pattern = "[^" + value + "]$"
			case "endswith":
				schema.Pattern = value + "$"
			case "excludes", "excludesall", "excludesrune":
				schema.Pattern = "^(?!" + value + ").*$"
			case "lowercase":
				schema.Pattern = "^[a-z]+$"
			case "multibyte":
				schema.Pattern = "[^\x00-\x7F]"
			case "number":
				schema.Type = "number"
			case "numeric":
				schema.Pattern = "^[0-9]+$"
			case "printascii":
				schema.Pattern = "^[\x20-\x7E]+$"
			case "startsnotwith":
				schema.Pattern = "^[^" + value + "]"
			case "startswith":
				schema.Pattern = "^" + value
			case "uppercase":
				schema.Pattern = "^[A-Z]+$"

			// String formats
			case "base64", "base64url":
				schema.Pattern = "^[A-Za-z0-9+/]*={0,2}$"
			case "base64rawurl":
				schema.Pattern = "^[A-Za-z0-9-_]*$"
			case "bic":
				schema.Pattern = "^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$"
			case "bcp47_language_tag":
				schema.Pattern = "^[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*$"
			case "btc_addr":
				schema.Pattern = "^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$"
			case "btc_addr_bech32":
				schema.Pattern = "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$"
			case "credit_card":
				schema.Pattern = "^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9]{2})[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35[0-9]{3})[0-9]{11})$"
			case "mongodb", "spicedb":
				schema.Pattern = "^[a-zA-Z0-9-_]+$"
			case "cron":
				schema.Pattern = "^((\\*(/[0-9]+)?|[0-9,-]+(,[0-9,-]+)*) ?){5}$"
			case "datetime":
				schema.Format = "date-time"
			case "e164":
				schema.Pattern = "^\\+[1-9]\\d{1,14}$"
			case "email":
				schema.Format = "email"
			case "eth_addr":
				schema.Pattern = "^0x[a-fA-F0-9]{40}$"
			case "hexadecimal":
				schema.Pattern = "^[a-fA-F0-9]+$"
			case "hexcolor":
				schema.Pattern = "^#?([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$"
			case "hsl":
				schema.Pattern = "^hsl\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*\\)$"
			case "hsla":
				schema.Pattern = "^hsla\\(\\s*\\d+(deg|grad|rad|turn)?(\\s*,\\s*\\d+%){2}\\s*,\\s*(0?.\\d+|1)\\s*\\)$"
			case "html":
				schema.Pattern = "^<([a-z]+)([^<]+)*(?:>(.*)<\\/\\1>|\\s+\\/>)$"
			case "html_encoded":
				schema.Pattern = "&([a-z]+|#[0-9]{1,6}|#x[0-9a-fA-F]{1,6});"
			case "isbn":
				schema.Pattern = "^(?:ISBN(?:-1[03])?:? )?(?=[-0-9 ]{17}$|[-0-9X ]{13}$|[0-9X]{10}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9X]$"
			case "isbn10":
				schema.Pattern = "^(?:ISBN(?:-10)?:? )?(?=[-0-9X ]{13}$|[0-9X]{10}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9X]$"
			case "isbn13":
				schema.Pattern = "^(?:ISBN(?:-13)?:? )?(?=[-0-9 ]{17}$|[0-9]{13}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?(?:[0-9]+[- ]?){2}[0-9]$"
			case "issn":
				schema.Pattern = "^\\d{4}-?\\d{3}[\\dX]$"
			case "iso3166_1_alpha2":
				schema.Pattern = "^[A-Z]{2}$"
			case "iso3166_1_alpha3":
				schema.Pattern = "^[A-Z]{3}$"
			case "iso3166_1_numeric":
				schema.Pattern = "^[0-9]{3}$"
			case "iso3166_2":
				schema.Pattern = "^[A-Z]{2}-[A-Z0-9]{1,3}$"
			case "iso4217":
				schema.Pattern = "^[A-Z]{3}$"
			case "jwt":
				schema.Pattern = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.([A-Za-z0-9-_]+)?$"
			case "latitude":
				schema.Pattern = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
			case "longitude":
				schema.Pattern = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
			case "luhn_checksum":
				schema.Pattern = "^[0-9]{12,19}$"
			case "postcode_iso3166_alpha2":
				schema.Pattern = "^[A-Z]{2}-[0-9]{5}$"
			case "postcode_iso3166_alpha2_field":
				schema.Pattern = "^[A-Z]{2}$"
			case "rgb":
				schema.Pattern = "^rgb\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*\\)$"
			case "rgba":
				schema.Pattern = "^rgba\\(\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*\\d+%?\\s*,\\s*(0?.\\d+|1)\\s*\\)$"
			case "ssn":
				schema.Pattern = "^(?!000|666)[0-8]\\d{2}-(?!00)\\d{2}-(?!0000)\\d{4}$"
			case "timezone":
				schema.Pattern = "^[a-zA-Z_]+/[a-zA-Z_]+$"
			case "uuid", "uuid3", "uuid3_rfc4122", "uuid4", "uuid4_rfc4122", "uuid5", "uuid5_rfc4122", "uuid_rfc4122":
				schema.Format = "uuid"
			case "md4", "md5":
				schema.Pattern = "^[a-fA-F0-9]{32}$"
			case "sha256":
				schema.Pattern = "^[a-fA-F0-9]{64}$"
			case "sha384":
				schema.Pattern = "^[a-fA-F0-9]{96}$"
			case "sha512":
				schema.Pattern = "^[a-fA-F0-9]{128}$"
			case "ripemd128", "ripemd160", "tiger128":
				schema.Pattern = "^[a-fA-F0-9]{32}$"
			case "tiger160", "tiger192":
				schema.Pattern = "^[a-fA-F0-9]{40}$"
			case "semver":
				schema.Pattern = "^v?(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$"
			case "ulid":
				schema.Pattern = "^[0-9A-Z]{26}$"
			case "cve":
				schema.Pattern = "^CVE-[0-9]{4}-[0-9]{4,}$"

			// Network formats
			case "cidr", "cidrv4":
				schema.Pattern = "^([0-9]{1,3}\\.){3}[0-9]{1,3}\\/[0-9]{1,2}$"
			case "cidrv6":
				schema.Pattern = "^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\/[0-9]{1,3}$"
			case "datauri":
				schema.Pattern = "^data:[a-z]+\\/[a-z0-9-+.]+(;[a-z-]+=[a-z0-9-]+)*;base64,[a-z0-9!$&',()*+,;=\\-._~:@/?%\\s]*$"
			case "fqdn", "hostname", "hostname_port", "hostname_rfc1123":
				schema.Format = "hostname"
			case "ip", "ip_addr", "ipv4", "ip4_addr":
				schema.Format = "ipv4"
			case "ipv6", "ip6_addr":
				schema.Format = "ipv6"
			case "mac":
				schema.Pattern = "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
			case "tcp4_addr":
				schema.Pattern = "^([0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]+$"
			case "tcp6_addr":
				schema.Pattern = "^\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]:[0-9]+$"
			case "tcp_addr":
				schema.Pattern = "^(([0-9]{1,3}\\.){3}[0-9]{1,3}|\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]):[0-9]+$"
			case "udp4_addr":
				schema.Pattern = "^([0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]+$"
			case "udp6_addr":
				schema.Pattern = "^\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]:[0-9]+$"
			case "udp_addr":
				schema.Pattern = "^(([0-9]{1,3}\\.){3}[0-9]{1,3}|\\[([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\\]):[0-9]+$"
			case "unix_addr":
				schema.Pattern = "^/.+/.+"
			case "uri", "url", "http_url":
				schema.Format = "uri"
			case "url_encoded":
				schema.Pattern = "^[a-zA-Z0-9-_.!~*'()]+$"
			case "urn_rfc2141":
				schema.Pattern = "^urn:[a-zA-Z0-9][a-zA-Z0-9-]{0,31}:([a-zA-Z0-9()+,\\-.:=@;$_!*'%/?#]|%[0-9a-fA-F]{2})+$"

			// Primitives
			case "required":
				levelRequired = true

			// Comparisons:
			case "eq":
				schema.Pattern = "^" + value + "$"
			case "eq_ignore_case":
				schema.Pattern = "(?i)^" + value + "$"
			case "gt":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Pattern = "^[0-9]+$"
				schema.Minimum = floatValue
			case "gte":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Pattern = "^[0-9]+$"
				schema.Minimum = floatValue
				schema.ExclusiveMinimum = false
			case "lt":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Pattern = "^[0-9]+$"
				schema.Maximum = floatValue
			case "lte":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Pattern = "^[0-9]+$"
				schema.Maximum = floatValue
				schema.ExclusiveMaximum = false
			case "ne":
				schema.Pattern = "^(?!" + value + ").*$"
			case "ne_ignore_case":
				schema.Pattern = "(?i)^(?!" + value + ").*$"

			// Other
			case "len":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					continue
				}

				schema.Minimum = float64(intValue)
				schema.Maximum = float64(intValue)
			case "max":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Maximum = floatValue
				schema.ExclusiveMaximum = true
			case "min":
				floatValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					continue
				}

				schema.Minimum = floatValue
				schema.ExclusiveMinimum = true
			case "oneof":
				strs := strings.Split(value, " ")

				schema.Enum = make([]interface{}, len(strs))
				for i, str := range strs {
					schema.Enum[i] = str
				}

			// Aliases
			case "iscolor":
				hexColor, _ := s.validateTagsToSchema("hexcolor")
				hsl, _ := s.validateTagsToSchema("hsl")
				hsla, _ := s.validateTagsToSchema("hsla")
				rgb, _ := s.validateTagsToSchema("rgb")
				rgba, _ := s.validateTagsToSchema("rgba")

				schema = Schema{
					OneOf: []Schema{
						hexColor,
						hsl,
						hsla,
						rgb,
						rgba,
					},
				}
			case "country_code":
				iso3166_1_alpha2, _ := s.validateTagsToSchema("iso3166_1_alpha2")
				iso3166_1_alpha3, _ := s.validateTagsToSchema("iso3166_1_alpha3")
				iso3166_1_numeric, _ := s.validateTagsToSchema("iso3166_1_numeric")

				schema = Schema{
					OneOf: []Schema{
						iso3166_1_alpha2,
						iso3166_1_alpha3,
						iso3166_1_numeric,
					},
				}
			}

			if !levelRequired {
				topLevel.AllOf = append(topLevel.AllOf, schema)
			} else {
				required = true
			}
		}
	}

	return
}
