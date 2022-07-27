package model

type InboxResponse struct {
	Viewer struct {
		Pk                           int64         `json:"pk"`
		Username                     string        `json:"username"`
		FullName                     string        `json:"full_name"`
		IsPrivate                    bool          `json:"is_private"`
		ProfilePicUrl                string        `json:"profile_pic_url"`
		ProfilePicId                 string        `json:"profile_pic_id"`
		IsVerified                   bool          `json:"is_verified"`
		HasAnonymousProfilePicture   bool          `json:"has_anonymous_profile_picture"`
		ReelAutoArchive              string        `json:"reel_auto_archive"`
		HasHighlightReels            bool          `json:"has_highlight_reels"`
		IsUsingUnifiedInboxForDirect bool          `json:"is_using_unified_inbox_for_direct"`
		BizUserInboxState            int           `json:"biz_user_inbox_state"`
		WaAddressable                bool          `json:"wa_addressable"`
		WaEligibility                int           `json:"wa_eligibility"`
		InteropMessagingUserFbid     string        `json:"interop_messaging_user_fbid"`
		AccountBadges                []interface{} `json:"account_badges"`
	} `json:"viewer"`
	Inbox struct {
		Threads []struct {
			HasOlder bool `json:"has_older"`
			HasNewer bool `json:"has_newer"`
			Pending  bool `json:"pending"`
			Items    []struct {
				ItemId                 string      `json:"item_id"`
				UserId                 int64       `json:"user_id"`
				Timestamp              int64       `json:"timestamp"`
				ItemType               string      `json:"item_type"`
				ClientContext          string      `json:"client_context"`
				ShowForwardAttribution bool        `json:"show_forward_attribution"`
				ForwardScore           interface{} `json:"forward_score"`
				IsShhMode              bool        `json:"is_shh_mode"`
				IsSentByViewer         bool        `json:"is_sent_by_viewer"`
				UqSeqId                int         `json:"uq_seq_id"`
				Text                   string      `json:"text"`
				TqSeqId                int         `json:"tq_seq_id"`
			} `json:"items"`
			Canonical  bool   `json:"canonical"`
			ThreadId   string `json:"thread_id"`
			ThreadV2Id string `json:"thread_v2_id"`
			Users      []struct {
				Pk               int64  `json:"pk"`
				Username         string `json:"username"`
				FullName         string `json:"full_name"`
				IsPrivate        bool   `json:"is_private"`
				ProfilePicUrl    string `json:"profile_pic_url"`
				ProfilePicId     string `json:"profile_pic_id"`
				FriendshipStatus struct {
					Following          bool `json:"following"`
					Blocking           bool `json:"blocking"`
					IsPrivate          bool `json:"is_private"`
					IncomingRequest    bool `json:"incoming_request"`
					OutgoingRequest    bool `json:"outgoing_request"`
					IsBestie           bool `json:"is_bestie"`
					IsRestricted       bool `json:"is_restricted"`
					ReachabilityStatus int  `json:"reachability_status"`
					IsFeedFavorite     bool `json:"is_feed_favorite"`
				} `json:"friendship_status"`
				IsVerified                   bool          `json:"is_verified"`
				HasAnonymousProfilePicture   bool          `json:"has_anonymous_profile_picture"`
				IsSupervisionFeaturesEnabled bool          `json:"is_supervision_features_enabled"`
				HasHighlightReels            bool          `json:"has_highlight_reels"`
				IsUsingUnifiedInboxForDirect bool          `json:"is_using_unified_inbox_for_direct"`
				BizUserInboxState            int           `json:"biz_user_inbox_state"`
				WaAddressable                bool          `json:"wa_addressable"`
				WaEligibility                int           `json:"wa_eligibility"`
				InteropMessagingUserFbid     interface{}   `json:"interop_messaging_user_fbid"`
				AccountBadges                []interface{} `json:"account_badges"`
				FbidV2                       string        `json:"fbid_v2"`
			} `json:"users"`
			ViewerId                      int64         `json:"viewer_id"`
			LastActivityAt                int64         `json:"last_activity_at"`
			Muted                         bool          `json:"muted"`
			VcMuted                       bool          `json:"vc_muted"`
			EncodedServerDataInfo         string        `json:"encoded_server_data_info"`
			AdminUserIds                  []interface{} `json:"admin_user_ids"`
			ApprovalRequiredForNewMembers bool          `json:"approval_required_for_new_members"`
			Archived                      bool          `json:"archived"`
			ThreadHasAudioOnlyCall        bool          `json:"thread_has_audio_only_call"`
			PendingUserIds                []interface{} `json:"pending_user_ids"`
			LastSeenAt                    struct {
				Field1 struct {
					Timestamp    string `json:"timestamp"`
					ItemId       string `json:"item_id"`
					ShhSeenState struct {
					} `json:"shh_seen_state"`
					CreatedAt interface{} `json:"created_at"`
				} `json:"54082540125,omitempty"`
				Field2 struct {
					Timestamp    string `json:"timestamp"`
					ItemId       string `json:"item_id"`
					ShhSeenState struct {
					} `json:"shh_seen_state"`
					CreatedAt string `json:"created_at"`
				} `json:"54245388792"`
			} `json:"last_seen_at"`
			RelevancyScore     int    `json:"relevancy_score"`
			RelevancyScoreExpr int    `json:"relevancy_score_expr"`
			OldestCursor       string `json:"oldest_cursor"`
			NewestCursor       string `json:"newest_cursor"`
			Inviter            struct {
				Pk                         int64         `json:"pk"`
				Username                   string        `json:"username"`
				FullName                   string        `json:"full_name"`
				IsPrivate                  bool          `json:"is_private"`
				ProfilePicUrl              string        `json:"profile_pic_url"`
				ProfilePicId               string        `json:"profile_pic_id"`
				IsVerified                 bool          `json:"is_verified"`
				HasAnonymousProfilePicture bool          `json:"has_anonymous_profile_picture"`
				ReachabilityStatus         int           `json:"reachability_status"`
				HasHighlightReels          bool          `json:"has_highlight_reels"`
				AccountBadges              []interface{} `json:"account_badges"`
				ReelAutoArchive            string        `json:"reel_auto_archive,omitempty"`
				AllowedCommenterType       string        `json:"allowed_commenter_type,omitempty"`
				InteropMessagingUserFbid   string        `json:"interop_messaging_user_fbid,omitempty"`
				FbidV2                     string        `json:"fbid_v2,omitempty"`
				LikedClipsCount            int           `json:"liked_clips_count,omitempty"`
			} `json:"inviter"`
			LabelItems        []interface{} `json:"label_items"`
			LastPermanentItem struct {
				ItemId                 string      `json:"item_id"`
				UserId                 int64       `json:"user_id"`
				Timestamp              int64       `json:"timestamp"`
				ItemType               string      `json:"item_type"`
				ClientContext          string      `json:"client_context"`
				ShowForwardAttribution bool        `json:"show_forward_attribution"`
				ForwardScore           interface{} `json:"forward_score"`
				IsShhMode              bool        `json:"is_shh_mode"`
				IsSentByViewer         bool        `json:"is_sent_by_viewer"`
				UqSeqId                int         `json:"uq_seq_id"`
				Text                   string      `json:"text"`
				TqSeqId                int         `json:"tq_seq_id"`
			} `json:"last_permanent_item"`
			Named                     bool          `json:"named"`
			NextCursor                string        `json:"next_cursor"`
			PrevCursor                string        `json:"prev_cursor"`
			ThreadTitle               string        `json:"thread_title"`
			LeftUsers                 []interface{} `json:"left_users"`
			Spam                      bool          `json:"spam"`
			BcPartnership             bool          `json:"bc_partnership"`
			MentionsMuted             bool          `json:"mentions_muted"`
			ThreadType                string        `json:"thread_type"`
			ThreadSubtype             int           `json:"thread_subtype"`
			ChatActivityMuted         bool          `json:"chat_activity_muted"`
			OutgoingChatActivityMuted bool          `json:"outgoing_chat_activity_muted"`
			ThreadHasDropIn           bool          `json:"thread_has_drop_in"`
			VideoCallId               interface{}   `json:"video_call_id"`
			ShhModeEnabled            bool          `json:"shh_mode_enabled"`
			ShhTogglerUserid          interface{}   `json:"shh_toggler_userid"`
			ShhReplayEnabled          bool          `json:"shh_replay_enabled"`
			IsGroup                   bool          `json:"is_group"`
			InputMode                 int           `json:"input_mode"`
			ReadState                 int           `json:"read_state"`
			AssignedAdminId           int           `json:"assigned_admin_id"`
			Folder                    int           `json:"folder"`
			LastNonSenderItemAt       int64         `json:"last_non_sender_item_at"`
			BusinessThreadFolder      int           `json:"business_thread_folder"`
			Theme                     struct {
				Id string `json:"id"`
			} `json:"theme"`
			ThemeData                        interface{} `json:"theme_data"`
			ThreadLabel                      int         `json:"thread_label"`
			MarkedAsUnread                   bool        `json:"marked_as_unread"`
			ThreadContextItems               interface{} `json:"thread_context_items"`
			IsCloseFriendThread              bool        `json:"is_close_friend_thread"`
			HasGroupsXacIneligibleUser       bool        `json:"has_groups_xac_ineligible_user"`
			ThreadImage                      interface{} `json:"thread_image"`
			IsXacThread                      bool        `json:"is_xac_thread"`
			IsTranslationEnabled             bool        `json:"is_translation_enabled"`
			TranslationBannerImpressionCount int         `json:"translation_banner_impression_count"`
			SystemFolder                     int         `json:"system_folder"`
			IsFanclubSubscriberThread        bool        `json:"is_fanclub_subscriber_thread"`
			JoinableGroupLink                string      `json:"joinable_group_link"`
			GroupLinkJoinableMode            int         `json:"group_link_joinable_mode"`
			SmartSuggestion                  interface{} `json:"smart_suggestion"`
			IsCreatorSubscriberThread        bool        `json:"is_creator_subscriber_thread"`
			CreatorSubscriberThreadResponse  interface{} `json:"creator_subscriber_thread_response"`
			RtcFeatureSetStr                 string      `json:"rtc_feature_set_str"`
			PersistentMenuIcebreakers        struct {
				PersistentIcebreakers []interface{} `json:"persistent_icebreakers"`
				AreDefaultIcebreakers bool          `json:"are_default_icebreakers"`
			} `json:"persistent_menu_icebreakers"`
			PublicChatMetadata struct {
				IsPublic                  bool `json:"is_public"`
				IsPinnableToViewerProfile bool `json:"is_pinnable_to_viewer_profile"`
				IsPinnedToViewerProfile   bool `json:"is_pinned_to_viewer_profile"`
			} `json:"public_chat_metadata"`
			ResponsivenessCategory interface{} `json:"responsiveness_category"`
		} `json:"threads"`
		HasOlder      bool  `json:"has_older"`
		UnseenCount   int   `json:"unseen_count"`
		UnseenCountTs int64 `json:"unseen_count_ts"`
		PrevCursor    struct {
			CursorTimestampSeconds int `json:"cursor_timestamp_seconds"`
			CursorRelevancyScore   int `json:"cursor_relevancy_score"`
			CursorThreadV2Id       int `json:"cursor_thread_v2_id"`
		} `json:"prev_cursor"`
		NextCursor struct {
			CursorTimestampSeconds string `json:"cursor_timestamp_seconds"`
			CursorRelevancyScore   string `json:"cursor_relevancy_score"`
			CursorThreadV2Id       string `json:"cursor_thread_v2_id"`
		} `json:"next_cursor"`
		BlendedInboxEnabled bool `json:"blended_inbox_enabled"`
	} `json:"inbox"`
	SeqId                 int    `json:"seq_id"`
	PendingRequestsTotal  int    `json:"pending_requests_total"`
	HasPendingTopRequests bool   `json:"has_pending_top_requests"`
	Status                string `json:"status"`
}

type MessageResponse struct {
	//私信发送返回内容

}
