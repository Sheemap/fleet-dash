// Definition of environment types. This allows typescript to know whats available.

interface ImportMetaEnv {
    readonly VITE_APP_TITLE: string
    readonly VITE_EVE_CLIENT_ID: string
    readonly VITE_EVE_REDIRECT_URI: string
    readonly VITE_EVE_SCOPES: string
    readonly VITE_FLEETDASH_CORE_SCHEME: string
    readonly VITE_FLEETDASH_CORE_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
