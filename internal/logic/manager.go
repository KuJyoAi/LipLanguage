package logic

//// UploadStandardVideo 上传标准视频
//func UploadStandardVideo(ctx *gin.Context) {
//	answer := ctx.PostForm("answer")
//	videoFile, _, err1 := ctx.AnswerQuestionRequest.FormFile("video")
//	lipFile, _, err2 := ctx.AnswerQuestionRequest.FormFile("lip_video")
//
//	if err1 != nil || err2 != nil {
//		logrus.Infof("[api.UploadStandardVideo] answer=%v err1=%v err2=%v",
//			answer, err1, err2)
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "参数错误",
//		})
//		return
//	}
//
//	videoData, err := io.ReadAll(videoFile)
//	if err != nil {
//		logrus.Errorf("[api.UploadStandardVideo] %v", err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "视频错误",
//		})
//		return
//	}
//
//	lipData, err := io.ReadAll(lipFile)
//	if err != nil {
//		logrus.Errorf("[api.UploadStandardVideo] %v", err)
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "视频错误",
//		})
//		return
//	}
//
//	videoId, lipId, err := service.UploadStandardVideo(answer, videoData, lipData)
//
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "上传失败",
//		})
//		return
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{
//			"msg":   "上传成功",
//			"video": videoId,
//			"lip":   lipId,
//		})
//		return
//	}
//}
