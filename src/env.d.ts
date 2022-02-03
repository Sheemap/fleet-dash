interface ImportMetaEnv {
    readonly VITE_APP_TITLE: string
    readonly VITE_EVE_CLIENT_ID: string
    readonly VITE_EVE_CLIENT_SECRET: string
    readonly VITE_EVE_REDIRECT_URI: string
    readonly VITE_EVE_SCOPES: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
