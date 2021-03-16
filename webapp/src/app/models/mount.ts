/**
 * App market
 * OpenAPI spec version: 1.0.0
 * Contact: info@menucha.de
 *
 */

/**
 * Mount
 */
export interface Mount { 
    /**
     * Mount type
     */
    type?: string;
    /**
     * Mount source
     */
    source?: string;
    /**
     * Mount destination
     */
    destination?: string;
    /**
     * Mount options
     */
    options?: string;
}
