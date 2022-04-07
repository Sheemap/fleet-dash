function addToRunningAvg(current: number, count: number, value: number): number {
    return current + ((value - current) / count) || 0;
}

function removeFromRunningAvg(current: number, count: number, value: number): number {
    if (count - 1  <= 0) {
        return 0;
    }

    return ((current * count) - value) / (count - 1) || 0;
}

export { addToRunningAvg, removeFromRunningAvg };