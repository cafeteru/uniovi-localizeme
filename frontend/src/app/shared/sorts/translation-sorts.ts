import { checkNotNullParams, sortStrings } from './sort-columns';
import { Translation } from '../../types/translation';

export function sortTranslationByVersion(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.version, b.version);
    return validParams === 0 ? a.version - b.version : validParams;
}

export function sortTranslationByStage(a: Translation, b: Translation): number {
    const validParams = checkNotNullParams(a.stage, b.stage);
    if (validParams !== 0) {
        return validParams;
    }
    const validNames = checkNotNullParams(a.stage.name, b.stage.name);
    return validNames === 0 ? sortStrings(a.stage.name, b.stage.name) : validNames;
}

export function sortTranslationByActive(a: Translation, b: Translation): number {
    return checkNotNullParams(a.active, b.active);
}
