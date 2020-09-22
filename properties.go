package calculus

import "time"

// workbookProperties represents the worksheet propertues and metadata.

type WorkbookProperties interface {
	Application() string
	SetApplication(string)
	Creator() string
	SetCreator(creator string)
	LastModifiedBy() string
	SetLastModifiedBy(lastModifiedBy string)
	Title() string
	SetTitle(title string)
	Subject() string
	SetSubject(subject string)
	Description() string
	SetDescription(description string)
	Keywords() string
	SetKeywords(keywords string)
	Category() string
	SetCategory(category string)
	Date1904() bool
	SetDate1904(date1904 bool)
	Created() time.Time
	SetCreated(created time.Time)
	Modified() time.Time
	SetModified(modified time.Time)
}

type workbookProperties struct {
	application    string
	creator        string
	lastModifiedBy string
	title          string
	subject        string
	description    string
	keywords       string
	category       string
	date1904       bool
	created        time.Time
	modified       time.Time
}

// Application retrieves the application name property.
func (p *workbookProperties) Application() string {
	return p.application
}

// SetApplication updates the application name property.
func (p *workbookProperties) SetApplication(application string) {
	p.application = application
}

// Creator retrieves the creator name property.
func (p *workbookProperties) Creator() string {
	return p.creator
}

// SetCreator updates the creator name property.
func (p *workbookProperties) SetCreator(creator string) {
	p.creator = creator
}

// LastModifiedBy retrieves the last modifier name property.
func (p *workbookProperties) LastModifiedBy() string {
	return p.lastModifiedBy
}

// SetLastModifiedBy updates the last modifier name property.
func (p *workbookProperties) SetLastModifiedBy(lastModifiedBy string) {
	p.lastModifiedBy = lastModifiedBy
}

// Title retrieves the document title property.
func (p *workbookProperties) Title() string {
	return p.title
}

// SetTitle updates the document title property.
func (p *workbookProperties) SetTitle(title string) {
	p.title = title
}

// Subject retrieves the document subject property.
func (p *workbookProperties) Subject() string {
	return p.subject
}

// SetSubject updates the document subject property.
func (p *workbookProperties) SetSubject(subject string) {
	p.subject = subject
}

// Description retrieves the document description property.
func (p *workbookProperties) Description() string {
	return p.description
}

// SetDescription updates the document description property.
func (p *workbookProperties) SetDescription(description string) {
	p.description = description
}

// Keywords retrieves the document keyworkds property.
func (p *workbookProperties) Keywords() string {
	return p.keywords
}

// SetKeywords updates the document keyworkds property.
func (p *workbookProperties) SetKeywords(keywords string) {
	p.keywords = keywords
}

// Category retrieves the document category property.
func (p *workbookProperties) Category() string {
	return p.category
}

// SetCategory updates the document category property.
func (p *workbookProperties) SetCategory(category string) {
	p.category = category
}

// Date1904 retrieves the date 1904 enabled property.
func (p *workbookProperties) Date1904() bool {
	return p.date1904
}

// SetDate1904 updates the date 1904 enabled property.
func (p *workbookProperties) SetDate1904(date1904 bool) {
	p.date1904 = date1904
}

// Created retrieves the creation date property.
func (p *workbookProperties) Created() time.Time {
	return p.created
}

// SetCreated updates the creation date property.
func (p *workbookProperties) SetCreated(created time.Time) {
	p.created = created
}

// Modified retrieves the modification date property.
func (p *workbookProperties) Modified() time.Time {
	return p.modified
}

// SetModified updates the modification date property.
func (p *workbookProperties) SetModified(modified time.Time) {
	p.modified = modified
}
