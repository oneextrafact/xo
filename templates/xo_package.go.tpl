// Package {{ .Package }} contains the types for schema '{{ schema .Schema }}'.
package {{ .Package }}

// Code generated. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"github.com/google/uuid"
)
