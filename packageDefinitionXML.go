package ngssp

import "encoding/xml"

type NgsspPackageEnvelope struct {
	XMLName      xml.Name     `xml:"Envelope"`
	NgsspPackage NgsspPackage `xml:"Body>NgsspPackages>NgsspPackage"`
}

type NgsspPackage struct {
	CoherenceKey           string `xml:"CoherenceKey"`
	NgsspPackageKey        string
	NgsspPackageName       string
	NgsspPackageDefinition NgsspPackageDefinition
}

type NgsspPackageDefinition struct {
	Name                  string
	Description           string
	Flags                 PackageDefinitionFlags
	StartPeriod           string
	EndPeriod             string
	Price                 int
	PriceComercial        int
	SubscriptionPeriod    string
	SubscriptionPeriodUOM string
	GracePeriod           int
	GracePeriodUOM        string
	RegistrationSID       string
	RegSubsAction         string
	Benefits              []PackageBenefit `xml:"Benefits>Benefit"`
	Renewal               PackageRenewal
	Messages              []Message           `xml:"Messages>Message"`
	Provisioning          PackageProvisioning `xml:"Provisioning"`
	Relationship          PackageRelation     `xml:"Relationship"`
}

type PackageDefinitionFlags struct {
	IsConsent        string
	IsActive         string
	IsGift           string
	IsExtra          string
	IsExtend         string
	IsReplaceEndDate string
}

type PackageRenewal struct {
	Flags         PackageRenewalFlags
	MaxCycle      int
	Amount        float64
	SID           string
	Period        string
	PeriodUOM     string
	Progression   string
	ActionFailure string
	Suspend       struct {
		Period string
		UOM    string
	}
	Reminders []PackageRenewalReminder `xml:"Reminders>Reminder"`
}

type PackageRenewalFlags struct {
	Auto    string
	Option  string
	Charge  string
	Suspend string
}

type PackageRenewalReminder struct {
	Type   string
	Period int
	UOM    string
}

type PackageBenefit struct {
	Name  string
	Value float64
	UOM   string
}

type Message struct {
	Case      string
	Templates []MessageTemplate `xml:"Templates>Template"`
}

type MessageTemplate struct {
	Type    string
	Content []struct {
		MessageType string `xml:"messageType,attr"`
		Value       string `xml:",chardata"`
	}
}

type PackageProvisioning struct {
	Sequence []PackageProvisioningSequence `xml:"Sequence>System"`
	Cases    []PackageProvisioiningCase    `xml:"Cases>Case"`
}

type PackageProvisioningSequence struct {
	Seq       int
	Name      string
	Operation string
}

type PackageProvisioiningCase struct {
	Type       string
	Definition []PackageProvisioningCaseDefinition
}

type PackageProvisioningCaseDefinition struct {
	System     string
	Operation  string
	Parameters []PackageProvisioningCaseDefinitionParameters `xml:"Parameters>Row"`
}

type PackageProvisioningCaseDefinitionParameters struct {
	Num    int                                             `xml:"num"`
	Params []PackageProvisioningCaseDefinitionParameterSeq `xml:"Param"`
}

type PackageProvisioningCaseDefinitionParameterSeq struct {
	Seq       int
	Parameter string
	Value     string
	Uom       string
}

type PackageRelation struct {
	Upgrades      []PackageRelationChange `xml:"Upgrades>Upgrade"`
	Downgrades    []PackageRelationChange `xml:"Downgrades>Downgrade"`
	Incompatibles []string                `xml:"Incompatibles>Incompatible"`
}

type PackageRelationChange struct {
	From string `xml:"From"`
}
