/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

module.exports = angular.module('trafficOps.private.admin.profiles.parameters', [])
	.config(function($stateProvider, $urlRouterProvider) {
		$stateProvider
			.state('trafficOps.private.admin.profiles.parameters', {
				url: '/{profileId}/parameters',
				views: {
					profilesContent: {
						templateUrl: 'common/modules/table/profileParameters/table.profileParameters.tpl.html',
						controller: 'TableProfileParametersController',
						resolve: {
							profile: function($stateParams, profileService) {
								return profileService.getProfile($stateParams.profileId);
							},
							profileParameters: function($stateParams, profileParameterService) {
								return profileParameterService.getProfileParameters($stateParams.profileId);
							}
						}
					}
				}
			})
		;
		$urlRouterProvider.otherwise('/');
	});
