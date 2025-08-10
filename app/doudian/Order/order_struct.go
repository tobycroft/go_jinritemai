package Order

type OrderRet struct {
	St   int    `json:"st"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		OrderId                 string      `json:"order_id"`
		ShopOrderId             string      `json:"shop_order_id"`
		OrderStatus             int         `json:"order_status"`
		UserId                  string      `json:"user_id"`
		NowTs                   int         `json:"now_ts"`
		PayType                 int         `json:"pay_type"`
		OrderType               int         `json:"order_type"`
		BType                   int         `json:"b_type"`
		CBiz                    int         `json:"c_biz"`
		Biz                     int         `json:"biz"`
		ReceiveType             int         `json:"receive_type"`
		EExpress                int         `json:"e_express"`
		Repeat                  int         `json:"repeat"`
		IsDup                   int         `json:"is_dup"`
		PreReceiveInfoExist     int         `json:"pre_receive_info_exist"`
		HasWriteOffRecord       bool        `json:"has_write_off_record"`
		IsAlreadyModifyAmount   int         `json:"is_already_modify_amount"`
		UserIsAuth              int         `json:"user_is_auth"`
		CanModifyAmount         int         `json:"can_modify_amount"`
		ChangeAddr              int         `json:"change_addr"`
		StoreName               string      `json:"store_name"`
		WaitShipCount           int         `json:"wait_ship_count"`
		ShippedCount            int         `json:"shipped_count"`
		ProductCount            int         `json:"product_count"`
		TotalPostAmount         int         `json:"total_post_amount"`
		TotalPayAmount          int         `json:"total_pay_amount"`
		PayAmount               int         `json:"pay_amount"`
		PostAmount              int         `json:"post_amount"`
		TotalTaxAmount          int         `json:"total_tax_amount"`
		TotalIncludeTaxAmount   int         `json:"total_include_tax_amount"`
		TotalExcludingTaxAmount int         `json:"total_excluding_tax_amount"`
		TotalGoodsAmount        int         `json:"total_goods_amount"`
		PromotionAmount         int         `json:"promotion_amount"`
		ModifyAmount            int         `json:"modify_amount"`
		ModifyPostAmount        int         `json:"modify_post_amount"`
		SkuModifyAmount         int         `json:"sku_modify_amount"`
		ShopReceiveAmount       int         `json:"shop_receive_amount"`
		PromotionPayAmount      int         `json:"promotion_pay_amount"`
		EnvelopePromotionAmount int         `json:"envelope_promotion_amount"`
		TotalTaxAmountDesc      string      `json:"total_tax_amount_desc"`
		ActualPayAmount         int         `json:"actual_pay_amount"`
		CreateTime              int         `json:"create_time"`
		ConfirmTime             int         `json:"confirm_time"`
		PayTime                 int         `json:"pay_time"`
		LogisticsTime           int         `json:"logistics_time"`
		ReceiptTime             int         `json:"receipt_time"`
		GroupTime               int         `json:"group_time"`
		ExpShipTime             int         `json:"exp_ship_time"`
		OrderTypeDesc           string      `json:"order_type_desc"`
		PayTypeDesc             string      `json:"pay_type_desc"`
		WriteOffDesc            interface{} `json:"write_off_desc"`
		BuyerWords              string      `json:"buyer_words"`
		Remark                  string      `json:"remark"`
		Star                    int         `json:"star"`
		UserNickname            string      `json:"user_nickname"`
		HasWriteOff             bool        `json:"has_write_off"`
		HasMore                 bool        `json:"has_more"`
		PreSaleDesc             string      `json:"pre_sale_desc"`
		ReceiveInfo             interface{} `json:"receive_info"`
		ReceiverInfo            struct {
			PostReceiver string `json:"post_receiver"`
			PostTel      string `json:"post_tel"`
			PostAddr     struct {
				Province struct {
					Name     string `json:"name"`
					Id       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"province"`
				City struct {
					Name     string `json:"name"`
					Id       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"city"`
				Town struct {
					Name     string `json:"name"`
					Id       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"town"`
				Street struct {
					Name     string `json:"name"`
					Id       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"street"`
				Detail string `json:"detail"`
			} `json:"post_addr"`
			CanView           int           `json:"can_view"`
			PostTelType       int           `json:"post_tel_type"`
			ExpireTime        int           `json:"expire_time"`
			IsShowEditAddress bool          `json:"is_show_edit_address"`
			CanPostpone       bool          `json:"can_postpone"`
			ExtensionNumber   string        `json:"extension_number"`
			PostTelMask       string        `json:"post_tel_mask"`
			AddressTag        []interface{} `json:"address_tag"`
			UserAccountInfos  interface{}   `json:"user_account_infos"`
			UiType            string        `json:"ui_type"`
			BuyerTelInfo      struct {
				PayTel          string `json:"pay_tel"`
				PayTelType      int    `json:"pay_tel_type"`
				ExpireTime      int    `json:"expire_time"`
				CanPostpone     bool   `json:"can_postpone"`
				PayTelMask      string `json:"pay_tel_mask"`
				UiType          string `json:"ui_type"`
				CanView         int    `json:"can_view"`
				ExtensionNumber string `json:"extension_number"`
				NickName        string `json:"nick_name"`
			} `json:"buyer_tel_info"`
			ExtraMap interface{} `json:"extra_map"`
		} `json:"receiver_info"`
		PolicyInfo      interface{} `json:"policy_info"`
		OrderStatusInfo struct {
			DeadLineTime                    int           `json:"dead_line_time"`
			OrderStatusIcon                 int           `json:"order_status_icon"`
			OrderStatusText                 string        `json:"order_status_text"`
			OrderStatusRemark               string        `json:"order_status_remark"`
			OrderStatusOverRemark           string        `json:"order_status_over_remark"`
			ShowRule                        bool          `json:"show_rule"`
			OrderStatusOverRemarkV2         string        `json:"order_status_over_remark_v2"`
			IsAppointmentShip               bool          `json:"is_appointment_ship"`
			AppointmentShipTime             int           `json:"appointment_ship_time"`
			AppointmentShipTimeStr          string        `json:"appointment_ship_time_str"`
			CutDownSecond                   int           `json:"cut_down_second"`
			OrderStatusRemarkV2             string        `json:"order_status_remark_v2"`
			ShipTimeChangeDesc              string        `json:"ship_time_change_desc"`
			ShipTimeChangeHover             string        `json:"ship_time_change_hover"`
			OrderStatusRemarkType           string        `json:"order_status_remark_type"`
			OrderStatusRemarkHover          string        `json:"order_status_remark_hover"`
			OrderExtraText                  string        `json:"order_extra_text"`
			OrderExtraTextHover             string        `json:"order_extra_text_hover"`
			OrderExtraTextType              string        `json:"order_extra_text_type"`
			DeadLineTimeV3                  int           `json:"dead_line_time_v3"`
			OrderStatusRemarkV3             string        `json:"order_status_remark_v3"`
			OrderStatusOverRemarkV3         string        `json:"order_status_over_remark_v3"`
			OrderStatusRemarkTypeV3         string        `json:"order_status_remark_type_v3"`
			OrderStatusShortRemark          string        `json:"order_status_short_remark"`
			OrderStatusRemarkBackgroundType string        `json:"order_status_remark_background_type"`
			OrderStatusRemarkTypeV2         string        `json:"order_status_remark_type_v2"`
			OrderStatusOverRemarkType       string        `json:"order_status_over_remark_type"`
			OrderStatusRemarkScene          string        `json:"order_status_remark_scene"`
			EarlyPromiseDeadlineTimeDelta   int           `json:"early_promise_deadline_time_delta"`
			EarlyPromiseRemark              string        `json:"early_promise_remark"`
			EarlyPromiseRemarkV2            string        `json:"early_promise_remark_v2"`
			EarlyPromiseOverRemark          string        `json:"early_promise_over_remark"`
			EarlyPromiseRemarkType          string        `json:"early_promise_remark_type"`
			EarlyPromiseRemarkTypeV2        string        `json:"early_promise_remark_type_v2"`
			EarlyPromiseOverRemarkType      string        `json:"early_promise_over_remark_type"`
			DeadLineTimeDelta               int           `json:"dead_line_time_delta"`
			EarlyPromiseDeadlineTime        int           `json:"early_promise_deadline_time"`
			CompareWith                     int           `json:"compare_with"`
			RecommendLogisticsCompanyList   []interface{} `json:"recommend_logistics_company_list"`
		} `json:"order_status_info"`
		OperationActions []string `json:"operation_actions"`
		ActionMap        struct {
			ContactBuyer struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
					ModalTitle        string      `json:"modal_title"`
					ModalContent      string      `json:"modal_content"`
					OkText            string      `json:"ok_text"`
					CancelText        string      `json:"cancel_text"`
					OkUrl             string      `json:"ok_url"`
					ApiParams         interface{} `json:"api_params"`
					ApiMethod         string      `json:"api_method"`
					BizActionApi      string      `json:"biz_action_api"`
					ActionOption1Text string      `json:"action_option1_text"`
					ActionOption1Url  string      `json:"action_option1_url"`
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUi     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionUrl    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"contactBuyer"`
			EditAddress struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
					ModalTitle        string      `json:"modal_title"`
					ModalContent      string      `json:"modal_content"`
					OkText            string      `json:"ok_text"`
					CancelText        string      `json:"cancel_text"`
					OkUrl             string      `json:"ok_url"`
					ApiParams         interface{} `json:"api_params"`
					ApiMethod         string      `json:"api_method"`
					BizActionApi      string      `json:"biz_action_api"`
					ActionOption1Text string      `json:"action_option1_text"`
					ActionOption1Url  string      `json:"action_option1_url"`
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUi     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionUrl    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"editAddress"`
			Remark struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
					ModalTitle        string      `json:"modal_title"`
					ModalContent      string      `json:"modal_content"`
					OkText            string      `json:"ok_text"`
					CancelText        string      `json:"cancel_text"`
					OkUrl             string      `json:"ok_url"`
					ApiParams         interface{} `json:"api_params"`
					ApiMethod         string      `json:"api_method"`
					BizActionApi      string      `json:"biz_action_api"`
					ActionOption1Text string      `json:"action_option1_text"`
					ActionOption1Url  string      `json:"action_option1_url"`
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUi     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionUrl    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"remark"`
			Ship struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
					ModalTitle        string      `json:"modal_title"`
					ModalContent      string      `json:"modal_content"`
					OkText            string      `json:"ok_text"`
					CancelText        string      `json:"cancel_text"`
					OkUrl             string      `json:"ok_url"`
					ApiParams         interface{} `json:"api_params"`
					ApiMethod         string      `json:"api_method"`
					BizActionApi      string      `json:"biz_action_api"`
					ActionOption1Text string      `json:"action_option1_text"`
					ActionOption1Url  string      `json:"action_option1_url"`
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUi     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionUrl    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"ship"`
			ViewOrder struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
					ModalTitle        string      `json:"modal_title"`
					ModalContent      string      `json:"modal_content"`
					OkText            string      `json:"ok_text"`
					CancelText        string      `json:"cancel_text"`
					OkUrl             string      `json:"ok_url"`
					ApiParams         interface{} `json:"api_params"`
					ApiMethod         string      `json:"api_method"`
					BizActionApi      string      `json:"biz_action_api"`
					ActionOption1Text string      `json:"action_option1_text"`
					ActionOption1Url  string      `json:"action_option1_url"`
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUi     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionUrl    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"viewOrder"`
		} `json:"action_map"`
		Button          interface{} `json:"button"`
		OrderBottomCard interface{} `json:"order_bottom_card"`
		ProductItem     []struct {
			ItemOrderId     string `json:"item_order_id"`
			OrderStatus     int    `json:"order_status"`
			OrderType       string `json:"order_type"`
			PreSaleType     int    `json:"pre_sale_type"`
			ProcessType     int    `json:"process_type"`
			Biz             int    `json:"biz"`
			PriceHasTaxType string `json:"price_has_tax_type"`
			CBiz            int    `json:"c_biz"`
			TotalAmount     int    `json:"total_amount"`
			PostAmount      int    `json:"post_amount"`
			TaxAmount       int    `json:"tax_amount"`
			PayAmount       int    `json:"pay_amount"`
			ProductId       string `json:"product_id"`
			SkuId           int    `json:"sku_id"`
			MerchantSkuCode string `json:"merchant_sku_code"`
			ProductName     string `json:"product_name"`
			Img             string `json:"img"`
			SkuSpec         []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"sku_spec"`
			ComboAmount         int         `json:"combo_amount"`
			ComboNum            int         `json:"combo_num"`
			ShippedCount        int         `json:"shipped_count"`
			SkuWeight           int         `json:"sku_weight"`
			ItemOrderStatusDesc string      `json:"item_order_status_desc"`
			PackageStatusDesc   string      `json:"package_status_desc"`
			AftersaleService    []string    `json:"aftersale_service"`
			PolicyInfo          interface{} `json:"policy_info"`
			AfterSaleInfo       struct {
				UrlAfterSaleNo                int    `json:"url_after_sale_no"`
				AfterSaleText                 string `json:"after_sale_text"`
				AfterSaleRemark               string `json:"after_sale_remark"`
				AfterSaleId                   string `json:"after_sale_id"`
				HasPreSale                    bool   `json:"has_pre_sale"`
				AftersaleStatusClassColorType int    `json:"aftersale_status_class_color_type"`
				AftersaleStatusClassString    string `json:"aftersale_status_class_string"`
				AfterSaleStatusRemark         string `json:"after_sale_status_remark"`
				AfterSaleStatusRemarkHover    string `json:"after_sale_status_remark_hover"`
			} `json:"after_sale_info"`
			Tags []struct {
				Key               string      `json:"key"`
				Text              string      `json:"text"`
				HoverText         string      `json:"hover_text"`
				TagType           string      `json:"tag_type"`
				HelpDoc           string      `json:"help_doc"`
				Icon              string      `json:"icon"`
				BelongTo          string      `json:"belong_to"`
				QuestionHoverText string      `json:"question_hover_text"`
				Position          string      `json:"position"`
				UrlConfigMaps     interface{} `json:"url_config_maps"`
			} `json:"tags"`
			PromotionDetail   interface{} `json:"promotion_detail"`
			SkuBundleSubSkus  interface{} `json:"sku_bundle_sub_skus"`
			ItemWarehouseId   string      `json:"item_warehouse_id"`
			ItemWarehouseName string      `json:"item_warehouse_name"`
			ActionMap         struct {
			} `json:"action_map"`
			RelationOrder     interface{} `json:"relation_order"`
			TradeType         int         `json:"trade_type"`
			AvailableFetchCnt int         `json:"available_fetch_cnt"`
			Properties        []struct {
				Key        string `json:"key"`
				Text       string `json:"text"`
				Action     string `json:"action"`
				ActionText string `json:"action_text"`
			} `json:"properties"`
			SkuCustomizationInfo       interface{}   `json:"sku_customization_info"`
			GivenProductType           string        `json:"given_product_type"`
			GivenProductActivityType   string        `json:"given_product_activity_type"`
			GivenProductActivityKey    string        `json:"given_product_activity_key"`
			ProcessStatus              int           `json:"process_status"`
			ItemOrderStatusRemark      string        `json:"item_order_status_remark"`
			ItemOrderStatusRemarkHover string        `json:"item_order_status_remark_hover"`
			SkuCarModelDesc            string        `json:"sku_car_model_desc"`
			PrivilegeInfoList          []interface{} `json:"privilege_info_list"`
			PackageWeight              int           `json:"package_weight"`
			LowPriceInfo               interface{}   `json:"low_price_info"`
			CampainDetail              interface{}   `json:"campain_detail"`
			RelationSkuOrderIds        []interface{} `json:"relation_sku_order_ids"`
			AccessorySkuInfo           interface{}   `json:"accessory_sku_info"`
			ProductCustomPropertites   interface{}   `json:"product_custom_propertites"`
			RelatedOrderInfo           interface{}   `json:"related_order_info"`
			SkuOrderStatusInfo         struct {
				DeadLineTime                    int           `json:"dead_line_time"`
				OrderStatusIcon                 int           `json:"order_status_icon"`
				OrderStatusText                 string        `json:"order_status_text"`
				OrderStatusRemark               string        `json:"order_status_remark"`
				OrderStatusOverRemark           string        `json:"order_status_over_remark"`
				ShowRule                        bool          `json:"show_rule"`
				OrderStatusOverRemarkV2         string        `json:"order_status_over_remark_v2"`
				IsAppointmentShip               bool          `json:"is_appointment_ship"`
				AppointmentShipTime             int           `json:"appointment_ship_time"`
				AppointmentShipTimeStr          string        `json:"appointment_ship_time_str"`
				CutDownSecond                   int           `json:"cut_down_second"`
				OrderStatusRemarkV2             string        `json:"order_status_remark_v2"`
				ShipTimeChangeDesc              string        `json:"ship_time_change_desc"`
				ShipTimeChangeHover             string        `json:"ship_time_change_hover"`
				OrderStatusRemarkType           string        `json:"order_status_remark_type"`
				OrderStatusRemarkHover          string        `json:"order_status_remark_hover"`
				OrderExtraText                  string        `json:"order_extra_text"`
				OrderExtraTextHover             string        `json:"order_extra_text_hover"`
				OrderExtraTextType              string        `json:"order_extra_text_type"`
				DeadLineTimeV3                  int           `json:"dead_line_time_v3"`
				OrderStatusRemarkV3             string        `json:"order_status_remark_v3"`
				OrderStatusOverRemarkV3         string        `json:"order_status_over_remark_v3"`
				OrderStatusRemarkTypeV3         string        `json:"order_status_remark_type_v3"`
				OrderStatusShortRemark          string        `json:"order_status_short_remark"`
				OrderStatusRemarkBackgroundType string        `json:"order_status_remark_background_type"`
				OrderStatusRemarkTypeV2         string        `json:"order_status_remark_type_v2"`
				OrderStatusOverRemarkType       string        `json:"order_status_over_remark_type"`
				OrderStatusRemarkScene          string        `json:"order_status_remark_scene"`
				EarlyPromiseDeadlineTimeDelta   int           `json:"early_promise_deadline_time_delta"`
				EarlyPromiseRemark              string        `json:"early_promise_remark"`
				EarlyPromiseRemarkV2            string        `json:"early_promise_remark_v2"`
				EarlyPromiseOverRemark          string        `json:"early_promise_over_remark"`
				EarlyPromiseRemarkType          string        `json:"early_promise_remark_type"`
				EarlyPromiseRemarkTypeV2        string        `json:"early_promise_remark_type_v2"`
				EarlyPromiseOverRemarkType      string        `json:"early_promise_over_remark_type"`
				DeadLineTimeDelta               int           `json:"dead_line_time_delta"`
				EarlyPromiseDeadlineTime        int           `json:"early_promise_deadline_time"`
				CompareWith                     int           `json:"compare_with"`
				RecommendLogisticsCompanyList   []interface{} `json:"recommend_logistics_company_list"`
			} `json:"sku_order_status_info"`
			SkuUrlMap struct {
				CommissionFirstOrderWithdraw struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"CommissionFirstOrderWithdraw"`
				CommissionWithdraw struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"CommissionWithdraw"`
				FullFreeWithdraw struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"FullFreeWithdraw"`
				SignInWithdraw struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"SignInWithdraw"`
				AccountChangeDetail struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"account_change_detail"`
				AfterSaleList struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"after_sale_list"`
				AheadEarlyPromiseRemark struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string `json:"text"`
						Placeholder []struct {
							Key   string `json:"key"`
							Type  string `json:"type"`
							Label string `json:"label"`
							Value string `json:"value"`
							Style struct {
							} `json:"style"`
							Text      string      `json:"text"`
							RichHover interface{} `json:"rich_hover"`
						} `json:"placeholder"`
						ModalTitle string `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"ahead_early_promise_remark"`
				AheadPeriodReceiveRemark struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"ahead_period_receive_remark"`
				AheadPeriodSendRemark struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"ahead_period_send_remark"`
				AheadStatusRemark struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string `json:"text"`
						Placeholder []struct {
							Key   string `json:"key"`
							Type  string `json:"type"`
							Label string `json:"label"`
							Value string `json:"value"`
							Style struct {
							} `json:"style"`
							Text      string      `json:"text"`
							RichHover interface{} `json:"rich_hover"`
						} `json:"placeholder"`
						ModalTitle string `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"ahead_status_remark"`
				AlreadyChangePromise struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Border       string `json:"border"`
							BorderRadius string `json:"borderRadius"`
							Color        string `json:"color"`
							Padding      string `json:"padding"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"already_change_promise"`
				AppealCenterUrl struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"appeal_center_url"`
				CheckNegotiation struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"check_negotiation"`
				CheckSupplyOrder struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"check_supply_order"`
				EarlySendRule struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"early_send_rule"`
				GiftInfoUrl struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"gift_info_url"`
				GiftReceiveHover struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"gift_receive_hover"`
				LateDeliveryHover struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"late_delivery_hover"`
				OrderStatusEndHover struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string        `json:"text"`
						Placeholder []interface{} `json:"placeholder"`
						ModalTitle  string        `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"order_status_end_hover"`
				PriorityDeliveryDetail struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"priority_delivery_detail"`
				PriorityDeliveryRemark struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover struct {
						Text        string `json:"text"`
						Placeholder []struct {
							Key   string `json:"key"`
							Type  string `json:"type"`
							Label string `json:"label"`
							Value string `json:"value"`
							Style struct {
							} `json:"style"`
							Text      string      `json:"text"`
							RichHover interface{} `json:"rich_hover"`
						} `json:"placeholder"`
						ModalTitle string `json:"modal_title"`
					} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"priority_delivery_remark"`
				PromiseModal struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color  string `json:"color"`
							Cursor string `json:"cursor"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"promise_modal"`
				QicUrl struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"qic_url"`
				QicUrlForDetail struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"qic_url_for_detail"`
				RechargeLink struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"recharge_link"`
				RemoteCostRule struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"remote_cost_rule"`
				SendRule struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"send_rule"`
				SupplyChainOrderRechargeLink struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"supply_chain_order_recharge_link"`
				ToReport struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"to_report"`
				UserAppointmentReceiptTimeDesc struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"user_appointment_receipt_time_desc"`
				UserAppointmentShipTimeDesc struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style struct {
							Color string `json:"color"`
						} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"user_appointment_ship_time_desc"`
				WuyouIcon struct {
					Text           string `json:"text"`
					Url            string `json:"url"`
					Type           string `json:"type"`
					ComponentProps struct {
						Style interface{} `json:"style"`
					} `json:"component_props"`
					RichHover  interface{} `json:"rich_hover"`
					HoverProps struct {
						OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
					} `json:"hover_props"`
				} `json:"wuyou_icon"`
			} `json:"sku_url_map"`
		} `json:"product_item"`
		ShopOrderTag    []interface{} `json:"shop_order_tag"`
		PayAmountDetail []struct {
			Name      string `json:"name"`
			Amount    string `json:"amount"`
			Hover     string `json:"hover"`
			AmountInt int    `json:"amount_int"`
		} `json:"pay_amount_detail"`
		WayBillUrl            string        `json:"way_bill_url"`
		CrossBorderSendType   int           `json:"cross_border_send_type"`
		OrderAmountCard       interface{}   `json:"order_amount_card"`
		PayAmountDesc         string        `json:"pay_amount_desc"`
		ShopReceiveAmountDesc string        `json:"shop_receive_amount_desc"`
		SerialNumbers         interface{}   `json:"serial_numbers"`
		AddressTag            []interface{} `json:"address_tag"`
		SupportDetail         interface{}   `json:"support_detail"`
		NeedSerialNumber      bool          `json:"need_serial_number"`
		BTypeDesc             string        `json:"b_type_desc"`
		CBizDesc              string        `json:"c_biz_desc"`
		PriceDetail           interface{}   `json:"price_detail"`
		PromotionDetail       interface{}   `json:"promotion_detail"`
		PayTypeDescHover      string        `json:"pay_type_desc_hover"`
		ManualOrderType       int           `json:"manual_order_type"`
		OrderIdForShow        string        `json:"order_id_for_show"`
		OrderTagStamp         []interface{} `json:"order_tag_stamp"`
		UrlMap                struct {
			CommissionFirstOrderWithdraw struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"CommissionFirstOrderWithdraw"`
			CommissionWithdraw struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"CommissionWithdraw"`
			FullFreeWithdraw struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"FullFreeWithdraw"`
			SignInWithdraw struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"SignInWithdraw"`
			AccountChangeDetail struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"account_change_detail"`
			AfterSaleList struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"after_sale_list"`
			AheadEarlyPromiseRemark struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color string `json:"color"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_early_promise_remark"`
			AheadPeriodReceiveRemark struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_period_receive_remark"`
			AheadPeriodSendRemark struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_period_send_remark"`
			AheadStatusRemark struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color string `json:"color"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_status_remark"`
			AlreadyChangePromise struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Border       string `json:"border"`
						BorderRadius string `json:"borderRadius"`
						Color        string `json:"color"`
						Padding      string `json:"padding"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"already_change_promise"`
			AppealCenterUrl struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"appeal_center_url"`
			CheckNegotiation struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"check_negotiation"`
			CheckSupplyOrder struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"check_supply_order"`
			EarlySendRule struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"early_send_rule"`
			GiftInfoUrl struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"gift_info_url"`
			GiftReceiveHover struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"gift_receive_hover"`
			ImmunityDetailInfoUrl struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps interface{} `json:"hover_props"`
			} `json:"immunity_detail_info_url"`
			PriorityDeliveryDetail struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"priority_delivery_detail"`
			PriorityDeliveryRemark struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color string `json:"color"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"priority_delivery_remark"`
			PromiseModal struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color  string `json:"color"`
						Cursor string `json:"cursor"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"promise_modal"`
			QicUrl struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"qic_url"`
			QicUrlForDetail struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"qic_url_for_detail"`
			RechargeLink struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"recharge_link"`
			RemoteCostRule struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"remote_cost_rule"`
			SendRule struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"send_rule"`
			SupplierShopName             interface{} `json:"supplier_shop_name"`
			SupplyChainOrderRechargeLink struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"supply_chain_order_recharge_link"`
			ToReport struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"to_report"`
			UserAppointmentReceiptTimeDesc struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color string `json:"color"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"user_appointment_receipt_time_desc"`
			UserAppointmentShipTimeDesc struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color string `json:"color"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"user_appointment_ship_time_desc"`
			WuyouIcon struct {
				Text           string `json:"text"`
				Url            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"wuyou_icon"`
		} `json:"url_map"`
		UserProfileTag []struct {
			Key               string      `json:"key"`
			Text              string      `json:"text"`
			HoverText         string      `json:"hover_text"`
			TagType           string      `json:"tag_type"`
			HelpDoc           string      `json:"help_doc"`
			Icon              string      `json:"icon"`
			BelongTo          string      `json:"belong_to"`
			QuestionHoverText string      `json:"question_hover_text"`
			Position          string      `json:"position"`
			UrlConfigMaps     interface{} `json:"url_config_maps"`
		} `json:"user_profile_tag"`
		SupermarketOrderSerialNo string      `json:"supermarket_order_serial_no"`
		DeliverName              string      `json:"deliver_name"`
		DeliverMobile            string      `json:"deliver_mobile"`
		ReceiptTimeFmt           string      `json:"receipt_time_fmt"`
		LogisticsStatus          string      `json:"logistics_status"`
		GreetWords               string      `json:"greet_words"`
		TransferReceiverInfo     interface{} `json:"transfer_receiver_info"`
		TotalProductCount        int         `json:"total_product_count"`
		TotalPrice               int         `json:"total_price"`
		LatestLogisticInfo       struct {
			TrackingNo             string      `json:"tracking_no"`
			CompanyCode            string      `json:"company_code"`
			Track                  string      `json:"track"`
			DeliverName            string      `json:"deliver_name"`
			DeliverMobile          string      `json:"deliver_mobile"`
			LoStatusText           string      `json:"lo_status_text"`
			Tags                   interface{} `json:"tags"`
			TrackDetailDescription string      `json:"track_detail_description"`
			ShowTrack              bool        `json:"show_track"`
			LoStatusTextType       string      `json:"lo_status_text_type"`
			DescriptionTextType    string      `json:"description_text_type"`
			LoStatusTextV2         string      `json:"lo_status_text_v2"`
			ShowCallDispatch       bool        `json:"show_call_dispatch"`
		} `json:"latest_logistic_info"`
		CreateTimeStr         string      `json:"create_time_str"`
		AmountDetailMap       interface{} `json:"amount_detail_map"`
		ExtraTag              interface{} `json:"extra_tag"`
		ShopPrivilegeInfoList interface{} `json:"shop_privilege_info_list"`
		GiftReceiveTimeStr    string      `json:"gift_receive_time_str"`
		RelateInfos           interface{} `json:"relate_infos"`
		BaseCard              interface{} `json:"base_card"`
		ReceiverCommon        interface{} `json:"receiver_common"`
		IsOrderInAbTest       struct {
		} `json:"is_order_in_ab_test"`
	} `json:"data"`
	Page     int `json:"page"`
	Size     int `json:"size"`
	Total    int `json:"total"`
	ExtraMap struct {
		GivenProductGray string `json:"given_product_gray"`
	} `json:"extra_map"`
}
