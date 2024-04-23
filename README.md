# Promotion System

## Overview
This document outlines the design for a promotion system targeting the first 100 users who register a new account on the Cake system per campaign. Upon their first login, these users will receive a 30% discount voucher applicable to mobile phone fee top-up transactions (money transfer from bank account) made via the Cake app.

The document also addresses scalability considerations to support at least 100,000 concurrent users.

## System Design

The promotion system has 3 main components:
* Promotion Service.
* Redis: store the counter of users who log in to the campaign.
* MongoDB: campaign and voucher database.

<img src="./asset/promotion.drawio.png">

### CampaignLogin

| Field       | Type   | Description        | 
|:------------|:-------|:-------------------| 
| UserID      | string | ID of the user     | 
| CampaignID  | string | ID of the campaign |
| CreatedDate | string | Time UTC +7        |

Index:
* campaign_idx (CampaignID, UserID): unique, compound

### Voucher

| Field       | Type     | Description                           | 
|:------------|:---------|:--------------------------------------| 
| _id         | ObjectID | ID of the voucher                     | 
| UserID      | string   | ID of the user                        | 
| Discount    | int      | Percent: 0 -> 100                     | 
| Description | string   |                                       |
| ExpireDate  | string   | Expire Date: Time UTC +7              |
| VoucherType | string   | Which flow the voucher can be applied |
| CreatedDate | string   | Time UTC +7                           |

Index:
* voucher_idx (UserID)



