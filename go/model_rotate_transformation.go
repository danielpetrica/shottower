/*
shottower
Copyright (C) 2022 Rémy Boulanouar

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * Shottower
 *
 * Shottower is the open source version of Shotstack which is a video, image and audio editing service that allows for the automated generation of videos, images and audio using JSON and a RESTful API.  You arrange and configure an edit and POST it to the API which will render your media and provide a file  location when complete.  For more details visit [shottower](https://github.com/DblK/shottower) or checkout our [getting started](https://shotstack.io/docs/guide/) documentation.  There are two main API's, one for editing and generating assets (Edit API) and one for managing hosted assets (Serve API).  The Edit API base URL is: <b>http://0.0.0.0:4000/{version}</b>  The Serve API base URL is: <b>http://0.0.0.0:4000/serve/{version}</b>
 *
 * API version: stage
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import "github.com/spf13/cast"

// RotateTransformation - Rotate a clip by the specified angle in degrees. Rotation origin is set based on the clips `position`.
type RotateTransformation struct {

	// The angle to rotate the clip. Can be 0 to 360, or 0 to -360. Using a positive number rotates the clip clockwise, negative numbers counter-clockwise.
	Angle int32 `json:"angle,omitempty"`
}

func NewRotateTransformation(m map[string]interface{}) *RotateTransformation {
	transform := &RotateTransformation{}

	if m["angle"] != nil {
		transform.Angle = cast.ToInt32(m["angle"].(float64))
	}
	return transform
}

func (s *RotateTransformation) checkEnumValues() error {
	if s.Angle < -360 || s.Angle > 360 {
		return &EnumError{Schema: "Soundtrack", Field: "Angle", Value: s.Angle}
	}

	return nil
}

// AssertRotateTransformationRequired checks if the required fields are not zero-ed
func AssertRotateTransformationRequired(obj RotateTransformation) error {
	if err := obj.checkEnumValues(); err != nil {
		return err
	}

	return nil
}

// AssertRecurseRotateTransformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RotateTransformation (e.g. [][]RotateTransformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRotateTransformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRotateTransformation, ok := obj.(RotateTransformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRotateTransformationRequired(aRotateTransformation)
	})
}
