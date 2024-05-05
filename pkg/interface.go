/**
 * @author CHUAN
 * @date 2024-05-05 9:19 am
 */

package pkg

// ReadIntDirectly
//
//	@Description: read int directly
//	@param data
//	@return int
func ReadIntDirectly(data int) int {
	return data
}

// ReadIntFromInterface
//
//	@Description: read int from interface
//	@param data
//	@return int
func ReadIntFromInterface(data interface{}) int {
	return data.(int)
}
