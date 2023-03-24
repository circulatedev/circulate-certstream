package main

type StringCompare struct {
	str      string
	callback func(*CertStreamMessage)
}

type CertStreamMessage struct {
	MessageType string `json:"message_type"`
	Data        struct {
		UpdateType string `json:"update_type"`
		LeafCert   struct {
			Subject struct {
				Aggregated string `json:"aggregated"`
				C          string `json:"C"`
				ST         string `json:"ST"`
				L          string `json:"L"`
				O          string `json:"O"`
				OU         string `json:"OU"`
				CN         string `json:"CN"`
			} `json:"subject"`
			Extensions struct {
				KeyUsage               string `json:"keyUsage"`
				ExtendedKeyUsage       string `json:"extendedKeyUsage"`
				BasicConstraints       string `json:"basicConstraints"`
				SubjectKeyIdentifier   string `json:"subjectKeyIdentifier"`
				AuthorityKeyIdentifier string `json:"authorityKeyIdentifier"`
				AuthorityInfoAccess    string `json:"authorityInfoAccess"`
				SubjectAltName         string `json:"subjectAltName"`
				CertificatePolicies    string `json:"certificatePolicies"`
				CRLDistributionPoints  string `json:"crlDistributionPoints"`
			} `json:"extensions"`
			NotBefore           float32  `json:"not_before"`
			NotAfter            float32  `json:"not_after"`
			AsDer               string   `json:"as_der"`
			AllDomains          []string `json:"all_domains"`
			ValidationLevel     string   `json:"validation_level"`
			ParentSpkiSubjectFp string   `json:"parent_spki_subject_fp"`
			SpkiSubjectFp       string   `json:"spki_subject_fp"`
			TBSSerialNumber     string   `json:"tbs_serial_number"`
			Version             int      `json:"version"`
			SignatureAlgorithm  string   `json:"signature_algorithm"`
			Issuer              struct {
				C                string `json:"C"`
				O                string `json:"O"`
				OU               string `json:"OU"`
				CN               string `json:"CN"`
				BusinessCategory string `json:"businessCategory"`
				Street           string `json:"street"`
				PostalCode       string `json:"postalCode"`
			} `json:"issuer"`
			SubjectRaw         []string `json:"subject_raw"`
			IssuerRaw          []string `json:"issuer_raw"`
			PublicKeyAlgorithm string   `json:"public_key_algorithm"`
			PublicKey          struct {
				Type     string `json:"type"`
				Asn1     string `json:"asn1"`
				AsDer    string `json:"as_der"`
				BitSize  int    `json:"bit_size"`
				Modulus  string `json:"modulus"`
				Exponent int    `json:"exponent"`
			} `json:"public_key"`
		} `json:"leaf_cert"`
		Chain     []interface{} `json:"chain"`
		CertIndex int           `json:"cert_index"`
		Seen      float32       `json:"seen"`
		Source    struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"source"`
	} `json:"data"`
}
