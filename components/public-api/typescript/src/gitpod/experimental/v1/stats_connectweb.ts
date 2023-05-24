/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=ts"
// @generated from file gitpod/experimental/v1/stats.proto (package gitpod.experimental.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {GetUserStatsRequest, GetUserStatsResponse} from "./stats_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service gitpod.experimental.v1.StatsService
 */
export const StatsService = {
  typeName: "gitpod.experimental.v1.StatsService",
  methods: {
    /**
     * Retrieves the current user stats
     *
     * @generated from rpc gitpod.experimental.v1.StatsService.GetUserStats
     */
    getUserStats: {
      name: "GetUserStats",
      I: GetUserStatsRequest,
      O: GetUserStatsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
