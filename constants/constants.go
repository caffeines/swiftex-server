package constants

// AdminRole is the role type of admin
type AdminRole string

const (
	AdminType    string = "Admin"
	MerchantType string = "Merchant"
	RiderType    string = "Rider"
)

const (
	ShopOwner     string = "Owner"
	ShopModerator string = "Moderator"
)

const (
	// SuperAdmin has all sorts of access
	SuperAdmin AdminRole = "Super Admin"
	// Admin has all access exepet admin creation
	Admin AdminRole = "Admin"
	// Moderator has all access except admin creation and payment
	Moderator AdminRole = "Moderator"
	// ZoneManager has zone centric access
	ZoneManager AdminRole = "Zone Manager"
	Rider       string    = "Rider"
)

var Roles = []AdminRole{SuperAdmin, Admin, Moderator, ZoneManager}

const (
	// Active status: user can access
	Active    string = "Active"
	Cancelled string = "Cancelled"
	Deactive  string = "Deactive"
	Created   string = "Created"
	Delivered string = "Delivered"
	Accepted  string = "Accepted"
	Assigned  string = "Assigned"
	Apporved  string = "Approved"
	Pending   string = "Pending"
	Declined  string = "Declined"
	InTransit string = "In Transit"
	Returned  string = "Returned"
)

var AllStatus = []string{Active, Deactive}

const (
	COD  string = "COD"
	PAID string = "PAID"
)

const (
	Role   string = "role"
	Phone  string = "phone"
	UserID string = "userId"
)

const TrackIDSize = 8

const (
	Express string = "Express"
	Regular string = "Regular"
)
