import {GroupInfo} from "../../Shared/Models";

/**
 * User configuration
 */
export interface UserConfig {
    joinedGroups: GroupInfo[]
    ExcludedRules: ExcludedRule[]
}

/**
 * Exclude rule for lessons
 */
export interface ExcludedRule {

}