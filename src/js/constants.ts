const TOTAL_OVER_TIME_WIDGET_ID = 'total-over-time-widget';
const TOTAL_SLIDER_WIDGET_ID = 'total-slider-widget';
const INCOMING_JAM_WIDGET_ID = 'incoming-jam-widget';
const OUTGOING_JAM_WIDGET_ID = 'outgoing-jam-widget';
const OVERVIEW_WIDGET_ID = 'overview-widget';

const EVENT_TYPES = [
    "IncomingDamageEvent",
    "OutgoingDamageEvent",
    "IncomingHullEvent",
    "OutgoingHullEvent",
    "IncomingArmorEvent",
    "OutgoingArmorEvent",
    "IncomingShieldEvent",
    "OutgoingShieldEvent",
    "IncomingCapacitorEvent",
    "OutgoingCapacitorEvent",
    "IncomingNeutEvent",
    "OutgoingNeutEvent",
    "IncomingNosEvent",
    "OutgoingNosEvent"
];

const MULTI_SELECT_STYLE = {
    container: 'relative mx-auto w-full flex items-center justify-end cursor-pointer border-gray-300 text-base leading-snug outline-none text-zinc-400 absolute color-zinc-100',
    containerDisabled: 'cursor-default bg-gray-100',
    containerOpen: 'rounded-b-none',
    containerOpenTop: 'rounded-t-none',
    singleLabel: 'flex items-center h-full max-w-full absolute left-0 top-0 pointer-events-none bg-transparent leading-snug pl-3.5 pr-16 box-border text-zinc-100',
    singleLabelText: 'overflow-ellipsis overflow-hidden block whitespace-nowrap max-w-full',
    multipleLabel: 'flex items-center h-full absolute left-0 top-0 pointer-events-none bg-transparent leading-snug pl-3.5 text-zinc-100',
    search: 'w-full absolute inset-0 outline-none focus:ring-0 appearance-none border-bottom border-0 text-base font-sans pl-3.5 bg-zinc-800 mt-0 w-1/2  px-0.5 border-0 border-b-2 border-zinc-500 text-center text-zinc-100 focus:text-left focus:border-zinc-300',
    placeholder: 'flex items-center h-full absolute left-0 top-0 pointer-events-none bg-transparent leading-snug pl-3.5 text-gray-500',
    caret: 'bg-multiselect-caret bg-center bg-no-repeat w-2.5 h-4 py-px box-content mr-3.5 relative z-10 opacity-40 flex-shrink-0 flex-grow-0 transition-transform transform pointer-events-none',
    caretOpen: 'rotate-180 pointer-events-auto',
    clear: 'pr-3.5 relative z-10 opacity-40 transition duration-300 flex-shrink-0 flex-grow-0 flex hover:opacity-80',
    clearIcon: 'bg-multiselect-remove bg-center bg-no-repeat w-2.5 h-4 py-px box-content inline-block',
    spinner: 'bg-multiselect-spinner bg-center bg-no-repeat w-4 h-4 z-10 mr-3.5 animate-spin flex-shrink-0 flex-grow-0',
    dropdown: 'max-h-60 absolute -left-px -right-px -bottom-0 transform translate-y-full border border-zinc-200 -mt-px overflow-y-scroll z-50 bg-zinc-600 flex flex-col rounded-b',
    dropdownTop: '-translate-y-full top-px bottom-auto flex-col-reverse rounded-b-none rounded-t',
    dropdownHidden: 'hidden',
    options: 'flex flex-col p-0 m-0 list-none',
    optionsTop: 'flex-col-reverse',
    option: 'flex items-center justify-start box-border text-left cursor-pointer text-base leading-snug py-2 px-3 text-zinc-200 bg-zinc-800',
    optionPointed: 'text-slate-200 bg-zinc-600',
    optionSelected: 'text-slate-200 bg-zinc-500',
    optionDisabled: 'text-gray-300 cursor-not-allowed',
    optionSelectedPointed: 'text-slate-200 bg-slate-600 opacity-90',
    optionSelectedDisabled: 'text-green-100 bg-green-500 bg-opacity-50 cursor-not-allowed',
    noOptions: 'py-2 px-3 text-gray-600 bg-white text-left',
    noResults: 'py-2 px-3 text-zinc-300 bg-zinc-800 text-left',
    fakeInput: 'bg-transparent absolute left-0 right-0 -bottom-px w-full h-px border-0 p-0 appearance-none outline-none text-transparent',
    spacer: 'h-9 py-px box-content',
};

export {
    TOTAL_OVER_TIME_WIDGET_ID,
    TOTAL_SLIDER_WIDGET_ID,
    INCOMING_JAM_WIDGET_ID,
    OUTGOING_JAM_WIDGET_ID,
    OVERVIEW_WIDGET_ID,
    EVENT_TYPES,
    MULTI_SELECT_STYLE,
};