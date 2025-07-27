package gpcm

var (
	WWFCMsgUnknownLoginError = WWFCErrorMessage{
		ErrorCode: 22000,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCへの ログイン中に\n" +
				"不明なエラー が発生しました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"An unknown error has occurred\n" +
				"while logging in to Retro WFC.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Ein unbekannter Fehler ist beim\n" +
				"Verbinden mit Retro WFC aufgetreten.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Un error desconocido ha ocurrido\n" +
				"al conectarse a Retro WFC.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangItalian: "" +
				"È stato riscontrato un errore sconosciuto\n" +
				"durante l'accesso alla Retro WFC.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Er is een onbekende fout opgetreden\n" +
				"tijdens het verbinden met Retro WFC.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"載入Retro WFC 時發生錯誤\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"Retro WFC에 연결 도중\n" +
				"알 수 없는 오류가 발생했습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Při přihlašování k Retro WFC\n" +
				"došlo k neznámé chybě.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Во время входа в Retro WFC\n" +
				"произошла ошибка.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Retro Rewind'a giriş yaparken\n" +
				"bilinmeyen bir hata oluştu.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFinnish: "" +
				"Tapahtui Tuntematon virhe\n" +
				"kirjautuessa sisään Retro WFC:hen.\n" +
				"\n" +
				"Virhekoodi: %[1]d",
			LangFrenchEU: "" +
				"Une erreur inconnue s'est produite\n" +
				"pendant la connexion à Retro WFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d",
			LangSpanishEU: "" +
				"Un error desconocido ha ocurrido\n" +
				"al conectarse a Retro WFC.\n" +
				"\n" +
				"Código de error: %[1]d",
			LangPortugueseEU: "" +
				"Ocorreu um erro desconhecido\n" +
				"ao conectar-se com o Retro WFC.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgDolphinSetupRequired = WWFCErrorMessage{
		ErrorCode: 22001,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"あなたは すでにBANされている\n" +
				"リークされたNANDを使用しています\n" +
				"DolphinのデフォルトのNANDにもどしてください\n" +
				"チケットを discord.gg/retrorewind で切ってください\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are using a leaked NAND which is already banned.\n" +
				"Please revert your NAND to the Dolphin Default.\n" +
				"Ticket on discord.gg/retrorewind for help.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du benutzt eine geleakte NAND, die gebannt wurde.\n" +
				"Bitte wechsle zurück auf die Standard Dolphin NAND.\n" +
				"Ticket über discord.gg/retrorewind für Hilfe erstellen.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Estás usando una NAND filtrada que ya está baneada.\n" +
				"Por favor usa la NAND predeterminada por Dolphin.\n" +
				"Ticket en discord.gg/retrorewind para recibir ayuda.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Stai utilizzando una NAND compromessa che è stata già bannata.\n" +
				"Reimposta la tua NAND a quella di default di Dolphin.\n" +
				"Fai un Ticket su discord.gg/retrorewind per chiedere aiuto.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je gebruikt een publieke NAND die al verbannen is.\n" +
				"Gebruik de standaard Dolphin-NAND.\n" +
				"Maak een ticket op discord.gg/retrorewind voor hulp.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangKorean: "" +
				"차단된 유출 NAND를 사용 중입니다.\n" +
				"Dolphin의 기본 NAND로 되돌리십시오.\n" +
				"도움이 필요한 경우 discord.gg/retrorewind 에서 티켓을 끊으십시오.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Používáš uniklý NAND, který je už zakázán.\n" +
				"Vrať prosím tvůj NAND do výchozího nastavení Dolphin.\n" +
				"Vstupenka na discord.gg/retrorewind pro pomoc.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вы используете слитую прошивку NAND, которая уже заблокирована.\n" +
				"Верните настройки NAND на Dolphin по умолчанию.\n" +
				"За помощью обращайтесь в поддержку discord.gg/retrorewind\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Kullandığınız NAND (sistem dosyaları) sızdırılmış.\n" +
				"Lütfen NAND'inizi Dolphin varsayılanına geri çevirin.\n" +
				"Yardım için discord.gg/retrorewind'dan destek alabilirsiniz.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous utilisez une NAND publique qui a déjà été bannie.\n" +
				"Veuillez retirer votre NAND de Dolphin.\n" +
				"Un ticket peut être fait sur discord.gg/retrorewind si besoin.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás usando una NAND filtrada que ya está baneada.\n" +
				"Por favor usa la NAND predeterminada por Dolphin.\n" +
				"Ticket en discord.gg/retrorewind para recibir ayuda.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Estás usando una NAND filtrada que ya está baneada.\n" +
				"Por favor usa la NAND predeterminada por Dolphin.\n" +
				"Ticket no discord.gg/retrorewind para receber ajuda.\n" +
				"\n" +
				"Código de Erro: %[1]d\n" +
				"Informacão de suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileBannedTOS = WWFCErrorMessage{
		ErrorCode: 22002,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"利用きやくに いはんしたため\n" +
				"Retro WFCから BANされました\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from Retro WFC\n" +
				"due to a violation of the\n" +
				"Terms of Service.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von Retro WFC\n" +
				"wegen eines Verstoßes der\n" +
				"Terms of Service gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de Retro WFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalla Retro WFC\n" +
				"a causa di una violazione dei\n" +
				"Termini di Servizio.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van Retro WFC\n" +
				"vanwege een overtreding van de\n" +
				"gebruiksvoorwaarden.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"由於違反服務條款\n" +
				"你禁止使用 Retro WFC\n" +
				"\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"이용약관 위반으로\n" +
				"Retro WFC 계정이\n" +
				"정지됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz Retro WFC\n" +
				"z důvodu porušení\n" +
				"Podmínek Služby.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено играть в Retro\n" +
				"WFC из-за нарушения условий\n" +
				"использования сервиса.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Hizmet Şartlarını ihlal ettiğinizden\n" +
				"dolayı Retro WFC'ye erişiminiz\n" +
				"yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) de Retro WFC\n" +
				"à cause d'une violation des\n" +
				"Conditions de Service.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de Retro WFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido do Retro WFC\n" +
				"devido a uma violação dos\n" +
				"Termos e Condições\n" +
				"\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileBannedTOSNow = WWFCErrorMessage{
		ErrorCode: 22002,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"利用きやくに いはんしたため\n" +
				"Retro WFCから BANされています\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from Retro WFC\n" +
				"due to a violation of the\n" +
				"Terms of Service.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von Retro WFC\n" +
				"wegen eines Verstoßes der\n" +
				"Terms of Service gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de Retro WFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalla\n" +
				"Retro WFC a causa di una violazione\n" +
				"dei Termini di Servizio.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van Retro WFC\n" +
				"vanwege een overtreding van de\n" +
				"gebruiksvoorwaarden.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"由於違反服務條款\n" +
				"你禁止使用 Retro WFC\n" +
				"\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"이용약관 위반으로\n" +
				"Retro WFC 계정이\n" +
				"정지됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz Retro WFC\n" +
				"z důvodu porušení\n" +
				"Podmínek Služby.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено играть\n" +
				"в Retro WFC из-за нарушения\n" +
				"условий использования сервиса.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Hizmet Şartlarını ihlal ettiğinizden\n" +
				"dolayı Retro WFC'ye erişiminiz\n" +
				"yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) de Retro WFC\n" +
				"à cause d'une violation des\n" +
				"Conditions de Service.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de Retro WFC\n" +
				"debido a una violación de los\n" +
				"Terminos de Servicio.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido do Retro WFC\n" +
				"devido a uma violação dos\n" +
				"Termos e Condições\n" +
				"\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestricted = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCの ルールにいはんしたため\n" +
				"オンライン対戦から BANされました\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public\n" +
				"matches due to a violation\n" +
				"of the Retro WFC Rules.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentl. Räumen\n" +
				"wegen eines Verstoßes der\n" +
				"Retro WFC Regeln gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Te han baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalle corse\n" +
				"pubbliche a causa di una violazione\n" +
				"delle regole della Retro WFC.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare\n" +
				"wedstrijden vanwege een overtreding\n" +
				"van de Retro WFC-regels.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"由於你違反 Retro WFC 規則\n" +
				"你已被禁止參加公開遊戲\n" +
				"\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"Retro WFC 규정 위반으로\n" +
				"공개 경기에서 차단됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných\n" +
				"zápasů z důvodu porušení\n" +
				"pravidel Retro WFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено участвовать\n" +
				"в публичных играх из-за\n" +
				"нарушения правил Retro WFC.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Retro WFC kurallarını ihlal\n" +
				"ettiğinizden dolayı herkese\n" +
				"açık yarışlara erişiminiz yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) des matchs\n" +
				"public à cause d'un violation d'une\n" +
				"des règles de Retro WFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido das partidas\n" +
				"públicas devido a uma violação\n" +
				"dos Termos e Condições\n" +
				"\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedNow = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCの ルールにいはんしたため\n" +
				"オンライン対戦から BANされています\n" +
				"\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You have been banned from public\n" +
				"matches due to a violation\n" +
				"of the Retro WFC Rules.\n" +
				"\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentl. Räumen\n" +
				"wegen eines Verstoßes der\n" +
				"Retro WFC Regeln gebannt.\n" +
				"\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x,",
			LangSpanish: "" +
				"Te han baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei stato bannato dalle corse\n" +
				"pubbliche a causa di una violazione\n" +
				"delle regole della Retro WFC.\n" +
				"\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare\n" +
				"wedstrijden vanwege een overtreding\n" +
				"van de Retro WFC-regels.\n" +
				"\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"由於你違反 Retro WFC 規則\n" +
				"你已無法參加公開遊戲\n" +
				"\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"Retro WFC 규정 위반으로\n" +
				"공개 경기에서 차단됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných\n" +
				"zápasů z důvodu porušení\n" +
				"pravidel Retro WFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено участвовать\n" +
				"в публичных играх из-за нарушения\n" +
				"правил Retro WFC.\n" +
				"\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Retro WFC kurallarını ihlal\n" +
				"ettiğinizden dolayı herkese\n" +
				"açık yarışlara erişiminiz yasaklanmıştır.\n" +
				"\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous avez été banni(e) des matchs\n" +
				"publics à cause d'un violation d'une\n" +
				"des règles de Retro WFC.\n" +
				"\n" +
				"Code Erreur:  %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de partidas públicas\n" +
				"debido a una violación de las\n" +
				"reglas de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d\n" +
				"Información de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido das partidas\n" +
				"públicas devido a uma violação\n" +
				"dos Termos e Condições\n" +
				"\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedCustom = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"オンライン対戦から BANされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public matches.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentlichen Matches ausgeschlossen.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei bannato dalle corse pubbliche.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare wedstrijden.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"您被禁止參加公開遊戲\n" +
				"原因：%[3]s\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"공개 경기에서 차단됐습니다.\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вам запрещено участвовать\n" +
				"в публичных играх.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Herkese açık yarışlara\n" +
				"erişiminiz yasaklanmıştır.\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous êtes banni(e) des matchs publics.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido das partidas públicas.\n" +
				"Razão: %[3]s\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgProfileRestrictedNowCustom = WWFCErrorMessage{
		ErrorCode: 22003,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"オンライン対戦から BANされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d\n" +
				"サポート情報： NG%08[2]x",
			LangEnglish: "" +
				"You are banned from public matches.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Support Info: NG%08[2]x",
			LangGerman: "" +
				"Du wurdest von öffentlichen Matches ausgeschlossen.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d\n" +
				"Support-Info: NG%08[2]x",
			LangSpanish: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangItalian: "" +
				"Sei bannato dalle corse pubbliche.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d\n" +
				"Supporto Informativo: NG%08[2]x",
			LangDutch: "" +
				"Je bent verbannen van openbare wedstrijden.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d\n" +
				"Ondersteuningsinformatie: NG%08[2]x",
			LangTradChinese: "" +
				"您被禁止參加公開遊戲\n" +
				"原因：%[3]s\n" +
				"錯誤代碼：%[1]d\n" +
				"支援資訊：NG%08[2]x",
			LangKorean: "" +
				"공개 경기에서 차단됐습니다.\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d\n" +
				"지원 정보: NG%08[2]x",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Отныне вам запрещено\n" +
				"участвовать в публичных играх.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d\n" +
				"Информация для поддержки: NG%08[2]x",
			LangTurkish: "" +
				"Herkese açık yarışlara\n" +
				"erişiminiz yasaklanmıştır.\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d\n" +
				"Destek Bilgisi: NG%08[2]x",
			LangFrenchEU: "" +
				"Vous êtes banni(e) des matchs publics.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d\n" +
				"Information Support: NG%08[2]x",
			LangSpanishEU: "" +
				"Estás baneado de partidas públicas.\n" +
				"Motivo: %[3]s\n" +
				"Código de Error: %[1]d\n" +
				"Info. de soporte: NG%08[2]x",
			LangPortugueseEU: "" +
				"Foste banido das partidas públicas.\n" +
				"Razão: %[3]s\n" +
				"Código de Erro: %[1]d\n" +
				"Informação de Suporte: NG%08[2]x",
		},
	}

	WWFCMsgKickedGeneric = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCから キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"Retro WFC.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest aus Retro WFC\n" +
				"gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Te han expulsado de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso\n" +
				"dalla Retro WFC.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit Retro WFC\n" +
				"geschopt.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"您已被踢出 Retro WFC\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"Retro WFC에서 추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi vyhozen z\n" +
				"Retro WFC.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Вас выгнали из Retro WFC.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Retro WFC'den atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"Retro WFC.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"Te han expulsado de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"Foste expulso do\n" +
				"Retro WFC.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgKickedModerator = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCの モデレーターから\n" +
				"キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"Retro WFC by a moderator.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest von einem Moderator\n" +
				"aus Retro WFC gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Un moderador te ha\n" +
				"expulsado de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso dalla\n" +
				"Retro WFC da un moderatore.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit Retro WFC\n" +
				"geschopt door een moderator.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"您已被房主踢出 Retro WFC\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"관리자에 의해 Retro WFC에서\n" +
				"추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi vyhozen z\n" +
				"Retro WFC moderátorem.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Модератор выгнал вас\n" +
				"из Retro WFC.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Bir moderatör tarafından\n" +
				"Retro WFC'den atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"Retro WFC par un modérateur.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"Un moderador te ha\n" +
				"expulsado de Retro WFC.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"Foste expulso do Retro WFC\n" +
				"por um moderador.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgKickedRoomHost = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"フレンドルームの ホストから\n" +
				"キックされました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from the\n" +
				"friend room by the room creator.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest von dem Raum-Ersteller\n" +
				"aus der Freundes-Lobby gekickt.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"El creador de la sala te ha\n" +
				"expulsado de ella.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso dalla\n" +
				"stanza dal suo creatore.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit de vriendenkamer\n" +
				"geschopt door de gastheer.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"您已被房主踢出好友房間\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"방장에 의해 추방됐습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Hostitel tě vyhodil\n" +
				"z místnosti.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Создатель группы выгнал вас.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Oda kurucusu tarafından\n" +
				"odadan atıldınız.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de la salle\n" +
				"par le créateur.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"El creador de la sala te ha\n" +
				"expulsado de ella.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"Foste expulso da sala\n" +
				"pelo criador da sala.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgKickedCustom = WWFCErrorMessage{
		ErrorCode: 22004,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCから キックされました\n" +
				"りゆう： %[3]s\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You have been kicked from\n" +
				"Retro WFC.\n" +
				"Reason: %[3]s\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Du wurdest aus Retro WFC\n" +
				"gekickt.\n" +
				"Grund: %[3]s\n" +
				"Fehlercode: %[1]d,",
			LangSpanish: "" +
				"Te han expulsado de Retro WFC.\n" +
				"Motivo: %[3]s\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato espulso\n" +
				"dalla Retro WFC.\n" +
				"Motivo: %[3]s\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je bent uit Retro WFC\n" +
				"geschopt.\n" +
				"Reden: %[3]s\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"您已被踢出 Retro WFC\n" +
				"\n" +
				"原因：%[3]s\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"Retro WFC에서 추방됐습니다.\n" +
				"\n" +
				"사유: %[3]s\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Máš zákaz veřejných zápasů.\n" +
				"Důvod: %[3]s\n" +
				"Kód Chyby: %[1]d\n" +
				"Informace o Podpoře: NG%08[2]x",
			LangRussian: "" +
				"Вас выгнали из Retro WFC.\n" +
				"\n" +
				"Причина: %[3]s\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Retro WFC'den atıldınız.\n" +
				"\n" +
				"Sebep: %[3]s\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été expulsé de\n" +
				"Retro WFC.\n" +
				"Raison: %[3]s\n" +
				"Error Code: %[1]d",
			LangSpanishEU: "" +
				"Te han expulsado de Retro WFC.\n" +
				"Motivo: %[3]s\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"Foste expulso do\n" +
				"Retro WFC.\n" +
				"Razão: %[3]s\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgConsoleMismatch = WWFCErrorMessage{
		ErrorCode: 22005,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"使われているコンソールは\n" +
				"このプロファイルが 登録されたときに\n" +
				"使われたコンソールでは ありません\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The console you are using is not\n" +
				"the device used to register this\n" +
				"profile.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Konsole, die du gerade nutzt,\n" +
				"ist nicht dieselbe, mit der dieses\n" +
				"Profil erstellt wurde.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"La console che stai usando non è\n" +
				"il dispositivo usato per\n" +
				"registrare questo profilo.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De console die je gebruikt is niet\n" +
				"het apparaat dat is gebruikt om dit\n" +
				"profiel te registreren.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"你正在使用的主機\n" +
				"並非用於註冊此用戶的裝置\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"사용중인 콘솔이 프로필에\n" +
				"등록된 기기 정보와 다릅니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Konzole, kterou používáš, není\n" +
				"zařízením používaným k registraci\n" +
				"tohoto profilu.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Ваш профиль был зарегистрирован\n" +
				"на другой консоли.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kullandığınız konsol, kaydolduğunuz\n" +
				"profille aynı değil.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"La console que vous utilisé\n" +
				"n'est pas l'appareil utilisé pour\n" +
				"enregistrer ce profil.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"A consola que estás a usar não é\n" +
				"o dispositivo que foi usado para registar este\n" +
				"perfil.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgConsoleMismatchDolphin = WWFCErrorMessage{
		ErrorCode: 22005,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"使われているコンソールは\n" +
				"このプロファイルが 登録されたときに\n" +
				"使われたコンソールでは ありません\n" +
				"NANDがただしく 設定されていることを\n" +
				"ご確認ください\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The console you are using is not\n" +
				"the device used to register this\n" +
				"profile. Please make sure you've\n" +
				"set up your NAND correctly.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Konsole, die du gerade nutzt,\n" +
				"ist nicht dieselbe, mit der dieses\n" +
				"Profil erstellt wurde. Bitte gehe sicher\n" +
				"dass du das NAND korrekt initialisiert hast.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil. Asegurate que has configurado\n" +
				"correctamente tu NAND.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"La console che stai usando non è\n" +
				"il dispositivo usato per registrare\n" +
				"questo profilo. Assicurati di avere\n" +
				"impostato la tua NAND correttamente.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De console die je gebruikt is niet\n" +
				"het apparaat dat is gebruikt om dit\n" +
				"profiel te registreren. Zorg ervoor dat de\n" +
				"NAND juist is ingesteld.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"你正在使用的主機\n" +
				"並非用於註冊此用戶的裝置\n" +
				"請確認已安裝正確的 NAND\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"사용중인 콘솔이 프로필에\n" +
				"등록된 기기가 아닙니다.\n" +
				"NAND가 제대로 설치 됐는지\n" +
				"확인해 주십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Konzole, kterou používáš, není\n" +
				"zařízením používaným k registraci\n" +
				"tohoto profilu. Ujisti se prosím, že\n" +
				"jsi správně nastavil NAND.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Ваш профиль был зарегистрирован\n" +
				"на другой консоли. Проверьте, что\n" +
				"память NAND настроена правильно.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kullandığınız konsol, kaydolduğunuz\n" +
				"profille aynı değil.\n" +
				"NAND'inizi doğru ayarladığınıza emin olun.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"La consonle que vous utilisé\n" +
				"n'est pas l'appareil utilisé pour\n" +
				"enregistrer ce profil. Assurez-vous d'avoir\n" +
				"configuré votre NAND correctement.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"La consola que estas usando no es el\n" +
				"dispositivo usado para registrar este\n" +
				"perfil. Asegurate que has configurado\n" +
				"correctamente tu NAND.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"A consola que estás a usar não é\n" +
				"o dispositivo que foi usado para registar este\n" +
				"perfil. Por favor verifica que configuraste\n" +
				"corretamente o teu NAND.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgProfileIDInvalid = WWFCErrorMessage{
		ErrorCode: 22006,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"あなたが登録しようとした\n" +
				"プロファイルIDは むこうです\n" +
				"新しくライセンスをつくりなおしてください\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The profile ID you are trying to\n" +
				"register is invalid.\n" +
				"Please create a new license.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Profil-ID, die du versuchst zu\n" +
				"registrieren, ist ungültig.\n" +
				"Bitte erstelle ein neues Profil.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"El perfil que está tratando de\n" +
				"registrar es invalido.\n" +
				"Cree una nueva licencia.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"L'ID del profilo che stai cercando\n" +
				"di registrare non è valido.\n" +
				"Crea una nuova patente e riprova.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Het profiel-ID dat je probeert te\n" +
				"registreren is ongeldig.\n" +
				"Maak een nieuw profiel aan.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"你註冊的資料 ID 無效\n" +
				"請建立新的授權\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"사용하시려는 프로필 ID는\n" +
				"등록하실 수 없습니다.\n" +
				"새로운 라이센스를 만드십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Profil, který se pokoušíš\n" +
				"zaregistrovat, je neplatný.\n" +
				"Vytvoř si prosím nový účet.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"ID профиля, который вы пытаетесь\n" +
				"зарегистрировать, — некорректный.\n" +
				"Создайте новое удостоверение.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kaydolmak istediğiniz profil ID'si\n" +
				"geçerli bir ID değil.\n" +
				"Lütfen yeni bir lisans oluşturun.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"L'ID du profil que vous essayez\n" +
				"d'enregistrer est invalide.\n" +
				"Veuillez créer une nouveau permis.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"El perfil que está tratando de\n" +
				"registrar es invalido.\n" +
				"Cree una nueva licencia.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"O perfil que estás a tentar\n" +
				"registar é inválido.\n" +
				"Por favor cria uma nova licensa.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgProfileIDInUse = WWFCErrorMessage{
		ErrorCode: 22007,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"あなたが登録しようとした\n" +
				"フレンドコードは すでに登録されています\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The friend code you are trying to\n" +
				"register is already in use.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Der Freundescode, den du gerade\n" +
				"registrierst, wird bereits verwendet.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"La clave de amigo que está tratando\n" +
				"de registrar, ya está en uso.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Il codice amico che stai cercando\n" +
				"di registrare è già in uso.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De vriendcode die je probeert te\n" +
				"registreren is al in gebruik.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"此好友代碼已被使用\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"해당 친구 코드는 이미 사용중입니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Kód kamaráda, který se pokoušíš\n" +
				"zaregistrovat, se už používá.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Код друга, который вы пытаетесь\n" +
				"зарегистрировать, уже занят.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Kaydolmak istediğiniz arkadaş kodu\n" +
				"zaten kullanımda.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Le code ami que vous essayez\n" +
				"d'enregistrer est déjà utilisé.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"La clave de amigo que está tratando\n" +
				"de registrar, ya está en uso.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"O código de amigo que estás a tentar\n" +
				"registar já está em uso.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgPayloadInvalid = WWFCErrorMessage{
		ErrorCode: 22008,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"Retro WFCの ペイロードがむこうです\n" +
				"ゲームを 再起動してください\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"The Retro WFC payload is invalid.\n" +
				"Try restarting your game.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Die Retro WFC Payload ist ungültig.\n" +
				"Versuche das Spiel neu zu starten.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Retro WFC no cargó correctamente\n" +
				"Intente reiniciar su juego.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Il payload della Retro WFC non è valido.\n" +
				"Prova a riavviare il gioco.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"De Retro WFC-payload is ongeldig.\n" +
				"Probeer het spel opnieuw op te starten.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"Retro WFC 負載失效\n" +
				"請嘗試重新啟動遊戲\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"Retro WFC 페이로드가 잘못됐습니다.\n" +
				"게임을 재시작 하십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Datový obsah Retro WFC je neplatné.\n" +
				"Zkus restartovat hru.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Запущен некорректный пейлоад\n" +
				"Retro WFC. Перезапустите игру.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Retro WFC payloud'u geçerli değil.\n" +
				"Oyunu yeniden başlatmayı deneyin.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Le payload Retro WFC est invalide.\n" +
				"Veuillez redémarrer votre jeu.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"Retro WFC no cargó correctamente\n" +
				"Intente reiniciar su juego.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"O RETRO WFC não carregou corretamente.\n" +
				"Tenta reiniciar o jogo.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgInvalidELO = WWFCErrorMessage{
		ErrorCode: 22009,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"VRまたはBRの値が むこうなため\n" +
				"Retro WFCから 切断されました\n" +
				"\n" +
				"エラーコード： %[1]d",
			LangEnglish: "" +
				"You were disconnected from\n" +
				"Retro WFC due to an invalid\n" +
				"VR or BR value.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Deine Verbindung zu Retro WFC\n" +
				"durch einen ungültigen VR oder BR\n" +
				"Wert beendet.\n" +
				"\n" +
				"Fehlercode: %[1]d",
			LangSpanish: "" +
				"Fuiste desconectado debido a discrepancias\n" +
				"con tu valor de PC o PB.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Sei stato disconnesso dalla Retro WFC\n" +
				"a causa di un valore non valido\n" +
				"di punti corsa o punti battaglia.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Je verbinding met Retro WFC is verbroken\n" +
				"vanwege een ongeldige rp- of gp-waarde.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"你與 Retro WFC 中斷連線\n" +
				"VR 或 BR 不會計算\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"잘못된 VR 또는 BR 값으로 인해\n" +
				"Retro WFC 연결이 끊어졌습니다.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Byl jsi odpojen od Retro WFC\n" +
				"z důvodu neplatné hodnoty\n" +
				"ZH nebo BH.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Вас отключили от Retro WFC\n" +
				"из-за некорректного значения\n" +
				"ГР или БР.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Hatalı bir KP veya SP\n" +
				"değerinden dolayı Retro WFC'ye\n" +
				"bağlantınız kesildi.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Vous avez été déconnecté de\n" +
				"Retro WFC à cause d'une valeur invalide\n" +
				"de Points Course ou Points Bataille.\n" +
				"\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"Fuiste desconectado debido a discrepancias\n" +
				"con tu valor de PC o PB.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"Foste desconectado do RETRO WFC\n" +
				"devido a um valor\n" +
				"CR ou CB inválido.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

	WWFCMsgInvalidHash = WWFCErrorMessage{
		ErrorCode: 22010,
		MessageRMC: map[byte]string{
			LangJapanese: "" +
				"むこうなパックのバージョンです！\n" +
				"ログインするために パックをアップデート\n" +
				"または インストールし直してください\n" +
				"\n" +
				"Error Code: %[1]d",
			LangEnglish: "" +
				"Invalid pack version!\n" +
				"Please update or reinstall your pack\n" +
				"to log in.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangGerman: "" +
				"Ungültige Pack-Version!\n" +
				"Bitte update oder installiere\n" +
				"das Pack neu, um dich einzuloggen.\n" +
				"\n" +
				"Error Code: %[1]d",
			LangSpanish: "" +
				"Versión del mod invalido.\n" +
				"Actualize o reinstale el pack para\n" +
				"poder ingresar.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangItalian: "" +
				"Versione della distribuzione errata!\n" +
				"Aggiorna o reinstalla la distribuzione\n" +
				"per connettesi.\n" +
				"\n" +
				"Codice Errore: %[1]d",
			LangDutch: "" +
				"Invalide versie van de mod!\n" +
				"Update of installeer de mod opnieuw\n" +
				"om in te loggen.\n" +
				"\n" +
				"Foutcode: %[1]d",
			LangTradChinese: "" +
				"套件版本不可用！\n" +
				"請更新或重新安裝套件\n" +
				"\n" +
				"錯誤代碼：%[1]d",
			LangKorean: "" +
				"팩 버전이 잘못됐습니다!\n" +
				"로그인을 하려면 팩을 업데이트\n" +
				"또는 재설치 하십시오.\n" +
				"\n" +
				"에러 코드: %[1]d",
			LangCzech: "" +
				"Neplatná verze balíčku!\n" +
				"Chceš-li se přihlásit, aktualizuj\n" +
				"nebo znovu nainstaluj balíček.\n" +
				"\n" +
				"Kód Chyby: %[1]d",
			LangRussian: "" +
				"Ваша версия RR устарела!\n" +
				"Чтобы войти, обновите\n" +
				"или переустановите RR.\n" +
				"\n" +
				"Код ошибки: %[1]d",
			LangTurkish: "" +
				"Geçersiz mod sürümü!\n" +
				"Oynayabilmek için lütfen mod'u\n" +
				"güncelleyin veya yeniden yükleyin.\n" +
				"\n" +
				"Hata Kodu: %[1]d",
			LangFrenchEU: "" +
				"Version du pack invalide!\n" +
				"Veuillez mettre à jour ou réinstaller\n" +
				"le pack pour se connecter.\n" +
				"Code Erreur: %[1]d",
			LangSpanishEU: "" +
				"Versión del mod invalido.\n" +
				"Actualize o reinstale el pack para\n" +
				"poder ingresar.\n" +
				"\n" +
				"Código de Error: %[1]d",
			LangPortugueseEU: "" +
				"A versão da distribuição é inválida!\n" +
				"Por favor atualiza ou reinstala a\n" +
				"distribuição para te conectares.\n" +
				"A.\n" +
				"\n" +
				"Código de Erro: %[1]d",
		},
	}

)