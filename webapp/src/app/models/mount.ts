/**
 * App market
 * OpenAPI spec version: 0.0.1
 * Contact: support@peraMIC.io
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
