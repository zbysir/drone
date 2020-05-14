// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fixtures

const PushHook = `
{
    "ref": "refs/heads/master",
    "before": "861f2315056e8925e627a6f46518b9df05896e24",
    "commits": [
        {
            "committer": {
                "name": "demo1",
                "email": "demo1@gmail.com"
            },
            "web_url": "https://coding.net/u/demo1/p/test1/git/commit/5b9912a6ff272e9c93a4c44c278fe9b359ed1ab4",
            "short_message": "new file .drone.yml\n",
            "sha": "5b9912a6ff272e9c93a4c44c278fe9b359ed1ab4"
        }
    ],
    "after": "5b9912a6ff272e9c93a4c44c278fe9b359ed1ab4",
    "event": "push",
    "repository": {
        "owner": {
            "path": "/u/demo1",
            "web_url": "https://coding.net/u/demo1",
            "global_key": "demo1",
            "name": "demo1",
            "avatar": "/static/fruit_avatar/Fruit-20.png"
        },
        "https_url": "https://git.coding.net/demo1/test1.git",
        "web_url": "https://coding.net/u/demo1/p/test1",
        "project_id": "99999999",
        "ssh_url": "git@git.coding.net:demo1/test1.git",
        "name": "test1",
        "description": ""
    },
    "user": {
        "path": "/u/demo1",
        "web_url": "https://coding.net/u/demo1",
        "global_key": "demo1",
        "name": "demo1",
        "avatar": "/static/fruit_avatar/Fruit-20.png"
    }
}
`

// 适配coding webhook v2
const PushHookV2 = `
{
  "ref": "refs/heads/master",
  "before": "cebbbb3f4bc0a4a04ca12fec5b094875f7bd0a5a",
  "after": "f2517d28df3ac6b033a4b8521b324ae5454c1a57",
  "created": false,
  "deleted": false,
  "compare": "https://xxx.coding.net/p/uc-api/git/compare/cebbbb3f4bc0a...f2517d28df3ac",
  "commits": [
    {
      "id": "f2517d28df3ac6b033a4b8521b324ae5454c1a57",
      "distinct": false,
      "message": "Accept Merge Request #1: (master-patch-1 -> master)\nMerge Request: ci\nCreated By: @张亮\nAccepted By: @张亮\nURL: https://xxx.coding.net/p/uc-api/d/uc-api/git/merge/1",
      "timestamp": 1589451706000,
      "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api/git/commit/f2517d28df3ac6b033a4b8521b324ae5454c1a57",
      "author": {
        "name": "张亮",
        "email": "zhangliang@zhuzi.me",
        "username": "张亮"
      },
      "committer": {
        "name": "张亮",
        "email": "zhangliang@zhuzi.me",
        "username": "张亮"
      },
      "added": [],
      "removed": [],
      "modified": [
        "README.md"
      ]
    },
    {
      "id": "6a639cb845c7e8828fa0e1004708bc121d757ef8",
      "distinct": false,
      "message": "更新文件 README.md\n",
      "timestamp": 1589451688000,
      "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api/git/commit/6a639cb845c7e8828fa0e1004708bc121d757ef8",
      "author": {
        "name": "张亮",
        "email": "zhangliang@zhuzi.me",
        "username": "张亮"
      },
      "committer": {
        "name": "张亮",
        "email": "zhangliang@zhuzi.me",
        "username": "张亮"
      },
      "added": [],
      "removed": [],
      "modified": [
        "README.md"
      ]
    }
  ],
  "head_commit": {
    "id": "f2517d28df3ac6b033a4b8521b324ae5454c1a57",
    "distinct": false,
    "message": "Accept Merge Request #1: (master-patch-1 -> master)\nMerge Request: ci\nCreated By: @张亮\nAccepted By: @张亮\nURL: https://xxx.coding.net/p/uc-api/d/uc-api/git/merge/1",
    "timestamp": 1589451706000,
    "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api/git/commit/f2517d28df3ac6b033a4b8521b324ae5454c1a57",
    "author": {
      "name": "张亮",
      "email": "zhangliang@zhuzi.me",
      "username": "张亮"
    },
    "committer": {
      "name": "张亮",
      "email": "zhangliang@zhuzi.me",
      "username": "张亮"
    },
    "added": [],
    "removed": [],
    "modified": [
      "README.md"
    ]
  },
  "pusher": {
    "name": "张亮",
    "email": "zhangliang@zhuzi.me",
    "username": "zhangliang"
  },
  "sender": {
    "id": 15271,
    "login": "zhangliang",
    "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-15.png",
    "url": "https://xxx.coding.net/api/user/key/zhangliang",
    "html_url": "https://xxx.coding.net/u/zhangliang",
    "name": "张亮",
    "name_pinyin": "zl|zhl|zhangliang"
  },
  "repository": {
    "id": 33774,
    "name": "uc-api",
    "full_name": "zhuzi/uc-api/uc-api",
    "owner": {
      "id": 21045,
      "login": "dongfeilong",
      "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
      "url": "https://xxx.coding.net/api/user/key/dongfeilong",
      "html_url": "https://xxx.coding.net/u/dongfeilong",
      "name": "董飞龙",
      "name_pinyin": "dfl|dongfeilong"
    },
    "private": true,
    "html_url": "https://xxx.coding.net/p/uc-api",
    "description": "用户中心api",
    "fork": false,
    "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api",
    "created_at": 1531397030000,
    "updated_at": 1567396699000,
    "clone_url": "https://e.coding.net/zhuzi/uc-api.git",
    "ssh_url": "git@e.coding.net:zhuzi/uc-api.git",
    "default_branch": "master"
  }
}
`

const DeleteBranchPushHook = `
{
    "ref": "refs/heads/master",
    "before": "861f2315056e8925e627a6f46518b9df05896e24",
    "after": "0000000000000000000000000000000000000000",
    "event": "push",
    "repository": {
        "owner": {
            "path": "/u/demo1",
            "web_url": "https://coding.net/u/demo1",
            "global_key": "demo1",
            "name": "demo1",
            "avatar": "/static/fruit_avatar/Fruit-20.png"
        },
        "https_url": "https://git.coding.net/demo1/test1.git",
        "web_url": "https://coding.net/u/demo1/p/test1",
        "project_id": "99999999",
        "ssh_url": "git@git.coding.net:demo1/test1.git",
        "name": "test1",
        "description": ""
    },
    "user": {
        "path": "/u/demo1",
        "web_url": "https://coding.net/u/demo1",
        "global_key": "demo1",
        "name": "demo1",
        "avatar": "/static/fruit_avatar/Fruit-20.png"
    }
}
`

const PullRequestHook = `
{
    "pull_request": {
        "target_branch": "master",
        "title": "pr1",
        "body": "pr message",
        "source_sha": "",
        "source_repository": {
            "owner": {
                "path": "/u/demo2",
                "web_url": "https://coding.net/u/demo2",
                "global_key": "demo2",
                "name": "demo2",
                "avatar": "/static/fruit_avatar/Fruit-2.png"
            },
            "https_url": "https://git.coding.net/demo2/test2.git",
            "web_url": "https://coding.net/u/demo2/p/test2",
            "project_id": "7777777",
            "ssh_url": "git@git.coding.net:demo2/test2.git",
            "name": "test2",
            "description": "",
            "git_url": "git://git.coding.net/demo2/test2.git"
        },
        "source_branch": "master",
        "number": 1,
        "web_url": "https://coding.net/u/demo1/p/test2/git/pull/1",
        "merge_commit_sha": "55e77b328b71d3ee4f9e70a5f67231b0acceeadc",
        "target_sha": "",
        "action": "create",
        "id": 7586,
        "user": {
            "path": "/u/demo2",
            "web_url": "https://coding.net/u/demo2",
            "global_key": "demo2",
            "name": "demo2",
            "avatar": "/static/fruit_avatar/Fruit-2.png"
        },
        "status": "CANMERGE"
    },
    "repository": {
        "owner": {
            "path": "/u/demo1",
            "web_url": "https://coding.net/u/demo1",
            "global_key": "demo1",
            "name": "demo1",
            "avatar": "/static/fruit_avatar/Fruit-20.png"
        },
        "https_url": "https://git.coding.net/demo1/test2.git",
        "web_url": "https://coding.net/u/demo1/p/test2",
        "project_id": "6666666",
        "ssh_url": "git@git.coding.net:demo1/test2.git",
        "name": "test2",
        "description": "",
        "git_url": "git://git.coding.net/demo1/test2.git"
    },
    "event": "pull_request"
}
`

const MergeRequestHook = `
{
    "merge_request": {
        "target_branch": "master",
        "title": "mr1",
        "body": "<p>mr message</p>",
        "source_sha": "",
        "source_branch": "branch1",
        "number": 1,
        "web_url": "https://coding.net/u/demo1/p/test1/git/merge/1",
        "merge_commit_sha": "74e6755580c34e9fd81dbcfcbd43ee5f30259436",
        "target_sha": "",
        "action": "create",
        "id": 533428,
        "user": {
            "path": "/u/demo1",
            "web_url": "https://coding.net/u/demo1",
            "global_key": "demo1",
            "name": "demo1",
            "avatar": "/static/fruit_avatar/Fruit-20.png"
        },
        "status": "CANMERGE"
    },
    "repository": {
        "owner": {
            "path": "/u/demo1",
            "web_url": "https://coding.net/u/demo1",
            "global_key": "demo1",
            "name": "demo1",
            "avatar": "/static/fruit_avatar/Fruit-20.png"
        },
        "https_url": "https://git.coding.net/demo1/test1.git",
        "web_url": "https://coding.net/u/demo1/p/test1",
        "project_id": "99999999",
        "ssh_url": "git@git.coding.net:demo1/test1.git",
        "name": "test1",
        "description": ""
    },
    "event": "merge_request"
}
`

const MergeRequestHookV2 = `
{
  "action": "merge",
  "number": 1,
  "mergeRequest": {
    "id": 6428648,
    "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api/git/merge/1",
    "html_url": "https://xxx.coding.net/p/uc-api/d/uc-api/git/merge/1",
    "patch_url": "https://xxx.coding.net/p/uc-api/d/uc-api/git/merge/1.patch",
    "diff_url": "https://xxx.coding.net/p/uc-api/d/uc-api/git/merge/1.diff",
    "number": 1,
    "state": "ACCEPTED",
    "title": "ci",
    "body": "",
    "user": {
      "id": 15271,
      "login": "zhangliang",
      "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-15.png",
      "url": "https://xxx.coding.net/api/user/key/zhangliang",
      "html_url": "https://xxx.coding.net/u/zhangliang",
      "name": "张亮",
      "name_pinyin": "zl|zhl|zhangliang"
    },
    "created_at": 1589451699000,
    "updated_at": 1589451699000,
    "merged_at": 1589451707067,
    "merge_commit_sha": "",
    "head": {
      "label": "zhuzi:master-patch-1",
      "ref": "master-patch-1",
      "sha": "6a639cb845c7e8828fa0e1004708bc121d757ef8",
      "user": {
        "id": 21045,
        "login": "dongfeilong",
        "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
        "url": "https://xxx.coding.net/api/user/key/dongfeilong",
        "html_url": "https://xxx.coding.net/u/dongfeilong",
        "name": "董飞龙",
        "name_pinyin": "dfl|dongfeilong"
      },
      "repo": {
        "id": 33774,
        "name": "uc-api",
        "full_name": "zhuzi/uc-api/uc-api",
        "owner": {
          "id": 21045,
          "login": "dongfeilong",
          "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
          "url": "https://xxx.coding.net/api/user/key/dongfeilong",
          "html_url": "https://xxx.coding.net/u/dongfeilong",
          "name": "董飞龙",
          "name_pinyin": "dfl|dongfeilong"
        },
        "private": true,
        "html_url": "https://xxx.coding.net/p/uc-api",
        "description": "用户中心api",
        "fork": false,
        "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api",
        "created_at": 1531397030000,
        "updated_at": 1567396699000,
        "clone_url": "https://e.coding.net/zhuzi/uc-api.git",
        "ssh_url": "git@e.coding.net:zhuzi/uc-api.git",
        "default_branch": "master"
      }
    },
    "base": {
      "label": "zhuzi:master",
      "ref": "master",
      "sha": "cebbbb3f4bc0a4a04ca12fec5b094875f7bd0a5a",
      "user": {
        "id": 21045,
        "login": "dongfeilong",
        "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
        "url": "https://xxx.coding.net/api/user/key/dongfeilong",
        "html_url": "https://xxx.coding.net/u/dongfeilong",
        "name": "董飞龙",
        "name_pinyin": "dfl|dongfeilong"
      },
      "repo": {
        "id": 33774,
        "name": "uc-api",
        "full_name": "zhuzi/uc-api/uc-api",
        "owner": {
          "id": 21045,
          "login": "dongfeilong",
          "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
          "url": "https://xxx.coding.net/api/user/key/dongfeilong",
          "html_url": "https://xxx.coding.net/u/dongfeilong",
          "name": "董飞龙",
          "name_pinyin": "dfl|dongfeilong"
        },
        "private": true,
        "html_url": "https://xxx.coding.net/p/uc-api",
        "description": "用户中心api",
        "fork": false,
        "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api",
        "created_at": 1531397030000,
        "updated_at": 1567396699000,
        "clone_url": "https://e.coding.net/zhuzi/uc-api.git",
        "ssh_url": "git@e.coding.net:zhuzi/uc-api.git",
        "default_branch": "master"
      }
    },
    "merged": true,
    "merged_by": {
      "id": 15271,
      "login": "zhangliang",
      "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-15.png",
      "url": "https://xxx.coding.net/api/user/key/zhangliang",
      "html_url": "https://xxx.coding.net/u/zhangliang",
      "name": "张亮",
      "name_pinyin": "zl|zhl|zhangliang"
    },
    "comments": 0,
    "commits": 1,
    "additions": 0,
    "deletions": 0,
    "changed_files": 1
  },
  "sender": {
    "id": 15271,
    "login": "zhangliang",
    "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-15.png",
    "url": "https://xxx.coding.net/api/user/key/zhangliang",
    "html_url": "https://xxx.coding.net/u/zhangliang",
    "name": "张亮",
    "name_pinyin": "zl|zhl|zhangliang"
  },
  "repository": {
    "id": 33774,
    "name": "uc-api",
    "full_name": "zhuzi/uc-api/uc-api",
    "owner": {
      "id": 21045,
      "login": "dongfeilong",
      "avatar_url": "https://coding.net/static/fruit_avatar/Fruit-4.png",
      "url": "https://xxx.coding.net/api/user/key/dongfeilong",
      "html_url": "https://xxx.coding.net/u/dongfeilong",
      "name": "董飞龙",
      "name_pinyin": "dfl|dongfeilong"
    },
    "private": true,
    "html_url": "https://xxx.coding.net/p/uc-api",
    "description": "用户中心api",
    "fork": false,
    "url": "https://xxx.coding.net/api/team/zhuzi/project/uc-api",
    "created_at": 1531397030000,
    "updated_at": 1567396699000,
    "clone_url": "https://e.coding.net/zhuzi/uc-api.git",
    "ssh_url": "git@e.coding.net:zhuzi/uc-api.git",
    "default_branch": "master"
  }
}
`
