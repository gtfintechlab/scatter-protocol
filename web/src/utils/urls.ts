function getBaseURL() {
    return "http://localhost:3000";
}

export const urls = {
    baseUrl: getBaseURL(),
    api: {
        auth: {
            signup: "/api/auth/sign-up",
            signin: "/api/auth/sign-in",
            logout: "/api/auth/logout",
            roles: "/api/auth/roles",
            displayName: "/api/auth/display-name",
            session: {
                general: "/api/auth/session"
            }
        },
        articles: {
            general: "/api/articles",
            single: "/api/articles/single",
            vote: "/api/articles/vote",
        },
        mailingList: {
            general: "/api/mailing-list"
        },
        comments: {
            general: "/api/comments"
        },
        admin: {
            users: {
                general: "/api/admin/users",
                promote: "/api/admin/users/promote",
                demote: "/api/admin/users/demote",
            },
            articles: {
                general: "/api/admin/articles",
                approve: "/api/admin/articles/approve",
                unapprove: "/api/admin/articles/unapprove",
                single: "/api/admin/articles/single",
            },
            comments: {
                general: "/api/admin/comments",
            }
        }
    },
    client: {
        home: "/",
        auth: {
            signup: "/sign-up",
            signin: "/sign-in",
        },
        articles: {
            single: "/article"
        }
    }
};