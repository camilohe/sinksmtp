// Provide a mapping from uint16 cipher IDs to names.
// The names are chosen to more or less match what Exim logs for no
// particularly strong reason.
//
// Name data is generated from 'openssl ciphers -V' because ha ha who
// wants to do this by hand? We assume that openssl on the author's
// system will cover everything that net/tls supports, which may be
// optimistic someday but is true so far.
// (Generated on Fedora 20 64-bit)

package main

// openssl ciphers -V | sort | field 1 3 4 7 | sed -e 's/,0x//' -e 's/Enc=.*(//' -e 's/)$//' | awk '{printf "%s: %s\n", $1, $3 ":" $2 ":" $4}' | sed -e 's/ / "/' -e 's/$/",/'
var cipherNames = map[uint16]string{
	// this will report something for uninitialized cipher values.
	0x0000: "NULL-CIPHER-ERROR",

	0x0004: "SSLv3:RC4-MD5:128",
	0x0005: "SSLv3:RC4-SHA:128",
	0x0007: "SSLv3:IDEA-CBC-SHA:128",
	0x000A: "SSLv3:DES-CBC3-SHA:168",
	0x0013: "SSLv3:EDH-DSS-DES-CBC3-SHA:168",
	0x0016: "SSLv3:EDH-RSA-DES-CBC3-SHA:168",
	0x001F: "SSLv3:KRB5-DES-CBC3-SHA:168",
	0x0020: "SSLv3:KRB5-RC4-SHA:128",
	0x0021: "SSLv3:KRB5-IDEA-CBC-SHA:128",
	0x0023: "SSLv3:KRB5-DES-CBC3-MD5:168",
	0x0024: "SSLv3:KRB5-RC4-MD5:128",
	0x0025: "SSLv3:KRB5-IDEA-CBC-MD5:128",
	0x002F: "SSLv3:AES128-SHA:128",
	0x0032: "SSLv3:DHE-DSS-AES128-SHA:128",
	0x0033: "SSLv3:DHE-RSA-AES128-SHA:128",
	0x0035: "SSLv3:AES256-SHA:256",
	0x0038: "SSLv3:DHE-DSS-AES256-SHA:256",
	0x0039: "SSLv3:DHE-RSA-AES256-SHA:256",
	0x003C: "TLSv1.2:AES128-SHA256:128",
	0x003D: "TLSv1.2:AES256-SHA256:256",
	0x0040: "TLSv1.2:DHE-DSS-AES128-SHA256:128",
	0x0041: "SSLv3:CAMELLIA128-SHA:128",
	0x0044: "SSLv3:DHE-DSS-CAMELLIA128-SHA:128",
	0x0045: "SSLv3:DHE-RSA-CAMELLIA128-SHA:128",
	0x0067: "TLSv1.2:DHE-RSA-AES128-SHA256:128",
	0x006A: "TLSv1.2:DHE-DSS-AES256-SHA256:256",
	0x006B: "TLSv1.2:DHE-RSA-AES256-SHA256:256",
	0x0084: "SSLv3:CAMELLIA256-SHA:256",
	0x0087: "SSLv3:DHE-DSS-CAMELLIA256-SHA:256",
	0x0088: "SSLv3:DHE-RSA-CAMELLIA256-SHA:256",
	0x008A: "SSLv3:PSK-RC4-SHA:128",
	0x008B: "SSLv3:PSK-3DES-EDE-CBC-SHA:168",
	0x008C: "SSLv3:PSK-AES128-CBC-SHA:128",
	0x008D: "SSLv3:PSK-AES256-CBC-SHA:256",
	0x0096: "SSLv3:SEED-SHA:128",
	0x0099: "SSLv3:DHE-DSS-SEED-SHA:128",
	0x009A: "SSLv3:DHE-RSA-SEED-SHA:128",
	0x009C: "TLSv1.2:AES128-GCM-SHA256:128",
	0x009D: "TLSv1.2:AES256-GCM-SHA384:256",
	0x009E: "TLSv1.2:DHE-RSA-AES128-GCM-SHA256:128",
	0x009F: "TLSv1.2:DHE-RSA-AES256-GCM-SHA384:256",
	0x00A2: "TLSv1.2:DHE-DSS-AES128-GCM-SHA256:128",
	0x00A3: "TLSv1.2:DHE-DSS-AES256-GCM-SHA384:256",
	0xC002: "SSLv3:ECDH-ECDSA-RC4-SHA:128",
	0xC003: "SSLv3:ECDH-ECDSA-DES-CBC3-SHA:168",
	0xC004: "SSLv3:ECDH-ECDSA-AES128-SHA:128",
	0xC005: "SSLv3:ECDH-ECDSA-AES256-SHA:256",
	0xC007: "SSLv3:ECDHE-ECDSA-RC4-SHA:128",
	0xC008: "SSLv3:ECDHE-ECDSA-DES-CBC3-SHA:168",
	0xC009: "SSLv3:ECDHE-ECDSA-AES128-SHA:128",
	0xC00A: "SSLv3:ECDHE-ECDSA-AES256-SHA:256",
	0xC00C: "SSLv3:ECDH-RSA-RC4-SHA:128",
	0xC00D: "SSLv3:ECDH-RSA-DES-CBC3-SHA:168",
	0xC00E: "SSLv3:ECDH-RSA-AES128-SHA:128",
	0xC00F: "SSLv3:ECDH-RSA-AES256-SHA:256",
	0xC011: "SSLv3:ECDHE-RSA-RC4-SHA:128",
	0xC012: "SSLv3:ECDHE-RSA-DES-CBC3-SHA:168",
	0xC013: "SSLv3:ECDHE-RSA-AES128-SHA:128",
	0xC014: "SSLv3:ECDHE-RSA-AES256-SHA:256",
	0xC023: "TLSv1.2:ECDHE-ECDSA-AES128-SHA256:128",
	0xC024: "TLSv1.2:ECDHE-ECDSA-AES256-SHA384:256",
	0xC025: "TLSv1.2:ECDH-ECDSA-AES128-SHA256:128",
	0xC026: "TLSv1.2:ECDH-ECDSA-AES256-SHA384:256",
	0xC027: "TLSv1.2:ECDHE-RSA-AES128-SHA256:128",
	0xC028: "TLSv1.2:ECDHE-RSA-AES256-SHA384:256",
	0xC029: "TLSv1.2:ECDH-RSA-AES128-SHA256:128",
	0xC02A: "TLSv1.2:ECDH-RSA-AES256-SHA384:256",
	0xC02B: "TLSv1.2:ECDHE-ECDSA-AES128-GCM-SHA256:128",
	0xC02C: "TLSv1.2:ECDHE-ECDSA-AES256-GCM-SHA384:256",
	0xC02D: "TLSv1.2:ECDH-ECDSA-AES128-GCM-SHA256:128",
	0xC02E: "TLSv1.2:ECDH-ECDSA-AES256-GCM-SHA384:256",
	0xC02F: "TLSv1.2:ECDHE-RSA-AES128-GCM-SHA256:128",
	0xC030: "TLSv1.2:ECDHE-RSA-AES256-GCM-SHA384:256",
	0xC031: "TLSv1.2:ECDH-RSA-AES128-GCM-SHA256:128",
	0xC032: "TLSv1.2:ECDH-RSA-AES256-GCM-SHA384:256",
}
