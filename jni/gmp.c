#include <string.h>
#include <stdio.h>
#include <jni.h>
#include "gme/gme.h"

static Music_Emu* emu;
static gme_info_t* info;

jobjectArray Java_com_mattjibson_gmp_GMFile_info(JNIEnv* env, jobject thiz, jstring fname)
{
	int tracks, i, len, cur;
	jobjectArray ret;
	jstring jstr;
	const char *str = (*env)->GetStringUTFChars(env, fname, 0);
	#define BUFSZ 16
	char play_len[BUFSZ];

	if(gme_open_file(str, &emu, gme_info_only))
	{
		(*env)->ReleaseStringUTFChars(env, fname, str);
		return (jobjectArray)(*env)->NewObjectArray(env, 0, (*env)->FindClass(env, "java/lang/String"), (*env)->NewStringUTF(env, ""));
	}

	tracks = gme_track_count(emu);

	(*env)->ReleaseStringUTFChars(env, fname, str);

	len = 3 + tracks * 2;

	ret = (jobjectArray)(*env)->NewObjectArray(env, len, (*env)->FindClass(env, "java/lang/String"), NULL);

	gme_track_info(emu, &info, 0);

	cur = 0;

	jstr = (*env)->NewStringUTF(env, info->system);
	(*env)->SetObjectArrayElement(env, ret, cur++, jstr);
	(*env)->DeleteLocalRef(env, jstr);

	jstr = (*env)->NewStringUTF(env, info->game);
	(*env)->SetObjectArrayElement(env, ret, cur++, jstr);
	(*env)->DeleteLocalRef(env, jstr);

	jstr = (*env)->NewStringUTF(env, info->author);
	(*env)->SetObjectArrayElement(env, ret, cur++, jstr);
	(*env)->DeleteLocalRef(env, jstr);

	gme_free_info(info);

	for(i = 0; i < tracks; i++)
	{
		gme_track_info(emu, &info, i);

		snprintf(play_len, BUFSZ, "%i", info->play_length);

		jstr = (*env)->NewStringUTF(env, info->song);
		(*env)->SetObjectArrayElement(env, ret, cur++, jstr);
		(*env)->DeleteLocalRef(env, jstr);

		jstr = (*env)->NewStringUTF(env, play_len);
		(*env)->SetObjectArrayElement(env, ret, cur++, jstr);
		(*env)->DeleteLocalRef(env, jstr);

		gme_free_info(info);
	}

	gme_delete(emu);

	return ret;
}
