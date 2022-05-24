import {GroupInfo} from "../../Shared/Models";

/**
 * User configuration
 */
export interface UserConfig {
    joinedGroups: GroupInfo[]
    excludedLessons: ExcludedRule[]
}

/**
 * Exclude rule for lessons
 */
export interface ExcludedRule {

}