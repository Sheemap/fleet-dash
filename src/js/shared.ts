function addToRunningAvg(current: number, count: number, value: number): number {
    return current + ((value - current) / count) || 0;
}

function removeFromRunningAvg(current: number, count: number, value: number): number {
    if (count - 1  <= 0) {
        return 0;
    }

    return ((current * count) - value) / (count - 1) || 0;
}

function createUUID(){
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random()*16|0, v = c == 'x' ? r : (r&0x3|0x8);
        return v.toString(16);
    });
}

class TotalCounterConfig {
    title: string;
    periodSeconds: number;
    eventTypes: string[];

    constructor(title: string, periodSeconds: number, eventTypes: string[]) {
        this.title = title;
        this.periodSeconds = periodSeconds;
        this.eventTypes = eventTypes;
    }
}

export type ValidModalConfigs = TotalCounterConfig;

export { addToRunningAvg, removeFromRunningAvg, createUUID, TotalCounterConfig };