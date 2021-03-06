/**
 * App market
 *
 * OpenAPI spec version: 1.0.0
 * Contact: info@menucha.de
 */
import { Device } from './device';
import { Mount } from './mount';

/**
 * App information
 */
export interface App {
    /**
     * App ID
     */
    id?: string;
    /**
     * App name
     */
    name?: string;
    /**
     * App label
     */
    label?: string;
    /**
     * Name of the app icon
     */
    icon?: string;
    /**
     * Location of the apps image file
     */
    image?: string;
    /**
     * Username for the repository access
     */
    user?: string;
    /**
     * Password for the repository access
     */
    passwd?: string;
    /**
     * Trust
     */
    trust?: boolean;
    /**
     * App state
     */
    state?: App.StateEnum;
    /**
     * Namespace
     */
    namespaces?: string;
    devices?: Array<Device>;
    mounts?: Array<Mount>;

    description?: string;
    inprogress?: boolean;
}
export namespace App {
    export type StateEnum = 'STARTING' | 'STARTED' | 'STOPPING' | 'STOPPED' | 'UPGRADINGSTARTED' | 'UPGRADINGSTOPPED' | 'RESETTINGSTARTED' | 'RESETTINGSTOPPED';
    export const StateEnum = {
      STARTING: 'STARTING' as StateEnum,
      STARTED: 'STARTED' as StateEnum,
      STOPPING: 'STOPPING' as StateEnum,
      STOPPED: 'STOPPED' as StateEnum,
      UPGRADINGSTARTED: 'UPGRADINGSTARTED' as StateEnum,
      UPGRADINGSTOPPED: 'UPGRADINGSTOPPED' as StateEnum,
      RESETTINGSTARTED: 'RESETTINGSTARTED' as StateEnum,
      RESETTINGSTOPPED: 'RESETTINGSTOPPED' as StateEnum
    };
}
